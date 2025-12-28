package router

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imkarthi24/sf-backend/internal/model/models"
	"github.com/imkarthi24/sf-backend/pkg/constants"
)

func RoleBasedAccessControl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionValue := ctx.Value(constants.SESSION)
		if sessionValue == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		session, ok := sessionValue.(*models.Session)
		if !ok || session == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Skip RBAC for system sessions (external access)
		if session.IsSystemSession {
			ctx.Next()
			return
		}

		if err := checkResourceAccess(ctx, session); err != nil {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.Next()
	}
}

// checkResourceAccess validates if the user has access to the requested resource
func checkResourceAccess(ctx *gin.Context, session *models.Session) error {
	// path := ctx.Request.URL.Path
	// method := ctx.Request.Method

	switch session.Role {

	default:
		return errors.New("unknown role: access denied")
	}
}

// checkStudentAccess validates student-specific access rules
func checkStudentAccess(ctx *gin.Context, session *models.Session, path, method string) error {
	// Check if the endpoint is in the blocked list
	if isStudentBlockedPath(path) {
		return errors.New("access denied: endpoint not allowed for students")
	}

	// Check if the endpoint is in the allowed list
	if !isStudentReadAllowed(path) {
		return errors.New("access denied: endpoint not allowed for students")
	}

	switch method {
	case "GET":
		if strings.Contains(path, "/api/sf/v1/student/") && !isStudentAccessingOwnData(ctx, session, path) {
			return errors.New("access denied: students can only access their own data")
		}
		if strings.Contains(path, "/api/sf/v1/entity-document/by-entity") && !isStudentAccessingOwnDocuments(ctx, session) {
			return errors.New("access denied: students can only access their own documents")
		}
		return nil

	case "PUT":
		if strings.Contains(path, "/api/sf/v1/student/") && !isStudentAccessingOwnData(ctx, session, path) {
			return errors.New("access denied: students can only modify their own data")
		}
		return nil

	case "POST":
		if strings.Contains(path, "/api/sf/v1/interview/") {
			return errors.New("access denied: students cannot create interview records")
		}
		if strings.Contains(path, "/api/sf/v1/placement/") {
			return errors.New("access denied: students cannot create placement records")
		}
		return nil

	case "DELETE":
		if strings.Contains(path, "/api/sf/v1/interview/") {
			return errors.New("access denied: students cannot delete interview records")
		}
		if strings.Contains(path, "/api/sf/v1/placement/") {
			return errors.New("access denied: students cannot delete placement records")
		}
		return nil

	default:
		return errors.New("access denied: method not allowed for students")
	}
}

func isStudentReadAllowed(path string) bool {
	// APIs that students CAN access (exact list from requirements)
	allowedPaths := []string{
		"/api/sf/v1/student/",                  // Student endpoints (with own data restriction)
		"/api/sf/v1/interview",                 // Interview management (full CRUD access)
		"/api/sf/v1/placement",                 // Placement management (full CRUD access)
		"/api/sf/v1/entity-document/by-entity", // Student's documents (checklist)
		"/api/sf/v1/course",                    // Course info (read-only)
		"/api/sf/v1/masterConfig/value",        // Master config values (read-only)
	}

	// Check if path is allowed
	for _, allowedPath := range allowedPaths {
		if strings.HasPrefix(path, allowedPath) {
			return true
		}
	}

	return false
}

// isStudentBlockedPath checks if a path is explicitly blocked for students
func isStudentBlockedPath(path string) bool {
	blockedPaths := []string{
		"/api/sf/v1/fee",            // Fee management
		"/api/sf/v1/feeDetail",      // Fee detail management
		"/api/sf/v1/enquiry",        // Enquiry management
		"/api/sf/v1/enquiryHistory", // Enquiry history
		"/api/sf/v1/candidate",      // Candidate management
		"/api/sf/v1/batch",          // Batch management
		"/api/sf/v1/entity/note",    // Entity notes
		"/api/sf/v1/entity/import",  // Entity import
		"/api/sf/v1/masterConfig",   // Master config management (except value)
		"/api/sf/v1/notification",   // Notification management
		"/api/sf/v1/telecaller",     // Telecaller management
		"/api/sf/v1/dashboard",      // Dashboard management
		"/api/sf/v1/user",           // User management (except login/reset)
	}

	for _, blockedPath := range blockedPaths {
		if strings.HasPrefix(path, blockedPath) {
			return true
		}
	}
	return false
}

// isStudentAccessingOwnData checks if a student is trying to access their own data
func isStudentAccessingOwnData(ctx *gin.Context, session *models.Session, path string) bool {
	// Handle empty student ID in session
	if session.StudentId == nil || *session.StudentId == 0 {
		return false
	}

	// Extract student ID from the path
	pathParts := strings.Split(path, "/")
	var studentIdFromPath string

	// Find the student ID in the path (after "/student/")
	// Pattern: /api/sf/v1/student/:id
	for i, part := range pathParts {
		if part == "student" && i+1 < len(pathParts) {
			studentIdFromPath = pathParts[i+1]
			// Additional validation: ensure the next part is not another path segment
			if studentIdFromPath != "" {
				break
			}
		}
	}

	if studentIdFromPath == "" {
		return false
	}

	// Convert to string for comparison
	return studentIdFromPath == fmt.Sprintf("%d", *session.StudentId)
}

// isStudentAccessingOwnInterviewData checks if a student is accessing their own interview data
func isStudentAccessingOwnInterviewData(ctx *gin.Context, session *models.Session, path string) bool {
	// Handle empty student ID in session
	if session.StudentId == nil || *session.StudentId == 0 {
		return false
	}

	// Path format: /api/sf/v1/interview/student/{studentId}
	pathParts := strings.Split(path, "/")

	// Find the student ID in the path (after "/student/")
	for i, part := range pathParts {
		if part == "student" && i+1 < len(pathParts) {
			studentIdFromPath := pathParts[i+1]
			return studentIdFromPath == fmt.Sprintf("%d", *session.StudentId)
		}
	}

	return false
}

// isStudentAccessingOwnPlacementData checks if a student is accessing their own placement data
func isStudentAccessingOwnPlacementData(ctx *gin.Context, session *models.Session, path string) bool {
	// Handle empty student ID in session
	if session.StudentId == nil || *session.StudentId == 0 {
		return false
	}

	// Path format: /api/sf/v1/placement/student/{studentId}
	pathParts := strings.Split(path, "/")

	// Find the student ID in the path (after "/student/")
	for i, part := range pathParts {
		if part == "student" && i+1 < len(pathParts) {
			studentIdFromPath := pathParts[i+1]
			return studentIdFromPath == fmt.Sprintf("%d", *session.StudentId)
		}
	}

	return false
}

// isStudentAccessingOwnDocuments checks if a student is accessing their own documents
func isStudentAccessingOwnDocuments(ctx *gin.Context, session *models.Session) bool {
	if session.StudentId == nil || *session.StudentId == 0 {
		return false
	}

	// Check query parameters: entityName=Student&entityId={studentId}
	entityName := ctx.Query("entityName")
	entityId := ctx.Query("entityId")

	if entityName != "Student" {
		return false
	}

	// Verify that the entityId matches the student's ID
	return entityId == fmt.Sprintf("%d", *session.StudentId)
}

// checkAttenderAccess handles access control for ATTENDER role
func checkAttenderAccess(ctx *gin.Context, session *models.Session, path, method string) error {
	// Implement specific rules for ATTENDER role
	// For now, allow basic access
	return nil
}

// checkTelecallerAccess handles access control for TELECALLER and TELECALLER_ADMIN roles
func checkTelecallerAccess(ctx *gin.Context, session *models.Session, path, method string) error {
	// Implement specific rules for TELECALLER roles
	// For now, allow basic access
	return nil
}

// checkExternalAccess handles access control for EXTERNAL role
func checkExternalAccess(ctx *gin.Context, session *models.Session, path, method string) error {
	// Implement specific rules for EXTERNAL role
	// For now, allow basic access
	return nil
}

// checkTrainerAccess handles access control for TRAINER role
func checkTrainerAccess(ctx *gin.Context, session *models.Session, path, method string) error {
	// Implement specific rules for TRAINER role
	// For now, allow basic access
	return nil
}
