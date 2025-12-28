package responseModel

type DashboardStats struct {
	EnquiryStats []StatusCountStat `json:"enquiryStats,omitempty"`
	StudentStats []StatusCountStat `json:"studentStats,omitempty"`
	FeeStats     *FeeStat          `json:"feeStats,omitempty"`

	AdmissionStats   *AdmissionStats   `json:"admissionStats,omitempty"`
	PlacementStats   *PlacementStats   `json:"placementStats,omitempty"`
	TelecallingStats *TelecallingStats `json:"telecallingStats,omitempty"`
	BatchStats       *BatchStats       `json:"batchStats,omitempty"`
	ReferrerStats    *ReferrerStats    `json:"referrerStats,omitempty"`
}

type StatusCountStat struct {
	Status string `json:"status,omitempty"`
	Count  int    `json:"count,omitempty"`
}

type FeeStat struct {
	TotalFeeAmount float64 `json:"totalFeeAmount,omitempty"`
	PaidAmount     float64 `json:"paidAmount,omitempty"`
	PendingAmount  float64 `json:"pendingAmount,omitempty"`
}

type AdmissionStats struct {
	SourceStats []SourceCountStat `json:"sourceStats,omitempty"`
	JoinedCount int               `json:"joinedCount,omitempty"`
}

type SourceCountStat struct {
	Source string `json:"source,omitempty"`
	Count  int    `json:"count,omitempty"`
}

type PlacementStats struct {
	OffersCount         int     `json:"offersCount,omitempty"`
	InterviewRatio      float64 `json:"interviewRatio,omitempty"`
	TotalInterviews     int     `json:"totalInterviews,omitempty"`
	SecondRoundsWaiting int     `json:"secondRoundsWaiting,omitempty"`
}

type TelecallingStats struct {
	CandidatesReached int `json:"candidatesReached,omitempty"`
	PositiveCount     int `json:"positiveCount,omitempty"`
	NegativeCount     int `json:"negativeCount,omitempty"`
	MovedToStudents   int `json:"movedToStudents,omitempty"`
}

type BatchStats struct {
	BatchesStarted         int               `json:"batchesStarted,omitempty"`
	BatchStatusStats       []StatusCountStat `json:"batchStatusStats,omitempty"`
	StudentAttendanceRatio float64           `json:"studentAttendanceRatio,omitempty"`
	StudentTestRatio       float64           `json:"studentTestRatio,omitempty"`
}

type ReferrerStats struct {
	ReferrerStats []ReferrerCountStat `json:"referrerStats,omitempty"`
}

type ReferrerCountStat struct {
	Referrer       string `json:"referrer,omitempty" gorm:"column:referrer"`
	ReferrerPhone  string `json:"referrerPhone,omitempty" gorm:"column:referrer_phone"`
	EnquiriesCount int    `json:"enquiriesCount,omitempty" gorm:"column:enquiries_count"`
	StudentsCount  int    `json:"studentsCount,omitempty" gorm:"column:students_count"`
}
