package entity

import "github.com/photoprism/photoprism/internal/classify"

const (
	// data sources
	SrcAuto     = ""
	SrcManual   = "manual"
	SrcEstimate = "estimate"
	SrcName     = "name"
	SrcMeta     = "meta"
	SrcXmp      = "xmp"
	SrcYaml     = "yaml"
	SrcLocation = classify.SrcLocation
	SrcImage    = classify.SrcImage

	// sort orders
	SortOrderRelevance = "relevance"
	SortOrderNewest    = "newest"
	SortOrderOldest    = "oldest"
	SortOrderImported  = "imported"
	SortOrderSimilar   = "similar"
	SortOrderName      = "name"

	// unknown values
	YearUnknown  = -1
	MonthUnknown = -1
	TitleUnknown = "Unknown"

	TypeDefault = ""
	TypeFolder  = "folder"
	TypeMoment  = "moment"
	TypeImage   = "image"
	TypeLive    = "live"
	TypeVideo   = "video"
	TypeRaw     = "raw"
	TypeText    = "text"

	RootDefault = ""
	RootImport  = "import"
	RootPath    = "/"
)
