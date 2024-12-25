// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/thomaspeugeot/phylotaxymusic/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	axisDBs map[uint]*AxisDB

	nextIDAxisDB uint

	axisgridDBs map[uint]*AxisGridDB

	nextIDAxisGridDB uint

	bezierDBs map[uint]*BezierDB

	nextIDBezierDB uint

	beziergridDBs map[uint]*BezierGridDB

	nextIDBezierGridDB uint

	beziergridstackDBs map[uint]*BezierGridStackDB

	nextIDBezierGridStackDB uint

	circleDBs map[uint]*CircleDB

	nextIDCircleDB uint

	circlegridDBs map[uint]*CircleGridDB

	nextIDCircleGridDB uint

	cursorDBs map[uint]*CursorDB

	nextIDCursorDB uint

	frontcurveDBs map[uint]*FrontCurveDB

	nextIDFrontCurveDB uint

	frontcurvestackDBs map[uint]*FrontCurveStackDB

	nextIDFrontCurveStackDB uint

	horizontalaxisDBs map[uint]*HorizontalAxisDB

	nextIDHorizontalAxisDB uint

	keyDBs map[uint]*KeyDB

	nextIDKeyDB uint

	parameterDBs map[uint]*ParameterDB

	nextIDParameterDB uint

	rhombusDBs map[uint]*RhombusDB

	nextIDRhombusDB uint

	rhombusgridDBs map[uint]*RhombusGridDB

	nextIDRhombusGridDB uint

	shapecategoryDBs map[uint]*ShapeCategoryDB

	nextIDShapeCategoryDB uint

	spiralbezierDBs map[uint]*SpiralBezierDB

	nextIDSpiralBezierDB uint

	spiralbeziergridDBs map[uint]*SpiralBezierGridDB

	nextIDSpiralBezierGridDB uint

	spiralcircleDBs map[uint]*SpiralCircleDB

	nextIDSpiralCircleDB uint

	spiralcirclegridDBs map[uint]*SpiralCircleGridDB

	nextIDSpiralCircleGridDB uint

	spirallineDBs map[uint]*SpiralLineDB

	nextIDSpiralLineDB uint

	spirallinegridDBs map[uint]*SpiralLineGridDB

	nextIDSpiralLineGridDB uint

	spiraloriginDBs map[uint]*SpiralOriginDB

	nextIDSpiralOriginDB uint

	spiralrhombusDBs map[uint]*SpiralRhombusDB

	nextIDSpiralRhombusDB uint

	spiralrhombusgridDBs map[uint]*SpiralRhombusGridDB

	nextIDSpiralRhombusGridDB uint

	verticalaxisDBs map[uint]*VerticalAxisDB

	nextIDVerticalAxisDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		axisDBs: make(map[uint]*AxisDB),

		axisgridDBs: make(map[uint]*AxisGridDB),

		bezierDBs: make(map[uint]*BezierDB),

		beziergridDBs: make(map[uint]*BezierGridDB),

		beziergridstackDBs: make(map[uint]*BezierGridStackDB),

		circleDBs: make(map[uint]*CircleDB),

		circlegridDBs: make(map[uint]*CircleGridDB),

		cursorDBs: make(map[uint]*CursorDB),

		frontcurveDBs: make(map[uint]*FrontCurveDB),

		frontcurvestackDBs: make(map[uint]*FrontCurveStackDB),

		horizontalaxisDBs: make(map[uint]*HorizontalAxisDB),

		keyDBs: make(map[uint]*KeyDB),

		parameterDBs: make(map[uint]*ParameterDB),

		rhombusDBs: make(map[uint]*RhombusDB),

		rhombusgridDBs: make(map[uint]*RhombusGridDB),

		shapecategoryDBs: make(map[uint]*ShapeCategoryDB),

		spiralbezierDBs: make(map[uint]*SpiralBezierDB),

		spiralbeziergridDBs: make(map[uint]*SpiralBezierGridDB),

		spiralcircleDBs: make(map[uint]*SpiralCircleDB),

		spiralcirclegridDBs: make(map[uint]*SpiralCircleGridDB),

		spirallineDBs: make(map[uint]*SpiralLineDB),

		spirallinegridDBs: make(map[uint]*SpiralLineGridDB),

		spiraloriginDBs: make(map[uint]*SpiralOriginDB),

		spiralrhombusDBs: make(map[uint]*SpiralRhombusDB),

		spiralrhombusgridDBs: make(map[uint]*SpiralRhombusGridDB),

		verticalaxisDBs: make(map[uint]*VerticalAxisDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *AxisDB:
		db.nextIDAxisDB++
		v.ID = db.nextIDAxisDB
		db.axisDBs[v.ID] = v
	case *AxisGridDB:
		db.nextIDAxisGridDB++
		v.ID = db.nextIDAxisGridDB
		db.axisgridDBs[v.ID] = v
	case *BezierDB:
		db.nextIDBezierDB++
		v.ID = db.nextIDBezierDB
		db.bezierDBs[v.ID] = v
	case *BezierGridDB:
		db.nextIDBezierGridDB++
		v.ID = db.nextIDBezierGridDB
		db.beziergridDBs[v.ID] = v
	case *BezierGridStackDB:
		db.nextIDBezierGridStackDB++
		v.ID = db.nextIDBezierGridStackDB
		db.beziergridstackDBs[v.ID] = v
	case *CircleDB:
		db.nextIDCircleDB++
		v.ID = db.nextIDCircleDB
		db.circleDBs[v.ID] = v
	case *CircleGridDB:
		db.nextIDCircleGridDB++
		v.ID = db.nextIDCircleGridDB
		db.circlegridDBs[v.ID] = v
	case *CursorDB:
		db.nextIDCursorDB++
		v.ID = db.nextIDCursorDB
		db.cursorDBs[v.ID] = v
	case *FrontCurveDB:
		db.nextIDFrontCurveDB++
		v.ID = db.nextIDFrontCurveDB
		db.frontcurveDBs[v.ID] = v
	case *FrontCurveStackDB:
		db.nextIDFrontCurveStackDB++
		v.ID = db.nextIDFrontCurveStackDB
		db.frontcurvestackDBs[v.ID] = v
	case *HorizontalAxisDB:
		db.nextIDHorizontalAxisDB++
		v.ID = db.nextIDHorizontalAxisDB
		db.horizontalaxisDBs[v.ID] = v
	case *KeyDB:
		db.nextIDKeyDB++
		v.ID = db.nextIDKeyDB
		db.keyDBs[v.ID] = v
	case *ParameterDB:
		db.nextIDParameterDB++
		v.ID = db.nextIDParameterDB
		db.parameterDBs[v.ID] = v
	case *RhombusDB:
		db.nextIDRhombusDB++
		v.ID = db.nextIDRhombusDB
		db.rhombusDBs[v.ID] = v
	case *RhombusGridDB:
		db.nextIDRhombusGridDB++
		v.ID = db.nextIDRhombusGridDB
		db.rhombusgridDBs[v.ID] = v
	case *ShapeCategoryDB:
		db.nextIDShapeCategoryDB++
		v.ID = db.nextIDShapeCategoryDB
		db.shapecategoryDBs[v.ID] = v
	case *SpiralBezierDB:
		db.nextIDSpiralBezierDB++
		v.ID = db.nextIDSpiralBezierDB
		db.spiralbezierDBs[v.ID] = v
	case *SpiralBezierGridDB:
		db.nextIDSpiralBezierGridDB++
		v.ID = db.nextIDSpiralBezierGridDB
		db.spiralbeziergridDBs[v.ID] = v
	case *SpiralCircleDB:
		db.nextIDSpiralCircleDB++
		v.ID = db.nextIDSpiralCircleDB
		db.spiralcircleDBs[v.ID] = v
	case *SpiralCircleGridDB:
		db.nextIDSpiralCircleGridDB++
		v.ID = db.nextIDSpiralCircleGridDB
		db.spiralcirclegridDBs[v.ID] = v
	case *SpiralLineDB:
		db.nextIDSpiralLineDB++
		v.ID = db.nextIDSpiralLineDB
		db.spirallineDBs[v.ID] = v
	case *SpiralLineGridDB:
		db.nextIDSpiralLineGridDB++
		v.ID = db.nextIDSpiralLineGridDB
		db.spirallinegridDBs[v.ID] = v
	case *SpiralOriginDB:
		db.nextIDSpiralOriginDB++
		v.ID = db.nextIDSpiralOriginDB
		db.spiraloriginDBs[v.ID] = v
	case *SpiralRhombusDB:
		db.nextIDSpiralRhombusDB++
		v.ID = db.nextIDSpiralRhombusDB
		db.spiralrhombusDBs[v.ID] = v
	case *SpiralRhombusGridDB:
		db.nextIDSpiralRhombusGridDB++
		v.ID = db.nextIDSpiralRhombusGridDB
		db.spiralrhombusgridDBs[v.ID] = v
	case *VerticalAxisDB:
		db.nextIDVerticalAxisDB++
		v.ID = db.nextIDVerticalAxisDB
		db.verticalaxisDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AxisDB:
		delete(db.axisDBs, v.ID)
	case *AxisGridDB:
		delete(db.axisgridDBs, v.ID)
	case *BezierDB:
		delete(db.bezierDBs, v.ID)
	case *BezierGridDB:
		delete(db.beziergridDBs, v.ID)
	case *BezierGridStackDB:
		delete(db.beziergridstackDBs, v.ID)
	case *CircleDB:
		delete(db.circleDBs, v.ID)
	case *CircleGridDB:
		delete(db.circlegridDBs, v.ID)
	case *CursorDB:
		delete(db.cursorDBs, v.ID)
	case *FrontCurveDB:
		delete(db.frontcurveDBs, v.ID)
	case *FrontCurveStackDB:
		delete(db.frontcurvestackDBs, v.ID)
	case *HorizontalAxisDB:
		delete(db.horizontalaxisDBs, v.ID)
	case *KeyDB:
		delete(db.keyDBs, v.ID)
	case *ParameterDB:
		delete(db.parameterDBs, v.ID)
	case *RhombusDB:
		delete(db.rhombusDBs, v.ID)
	case *RhombusGridDB:
		delete(db.rhombusgridDBs, v.ID)
	case *ShapeCategoryDB:
		delete(db.shapecategoryDBs, v.ID)
	case *SpiralBezierDB:
		delete(db.spiralbezierDBs, v.ID)
	case *SpiralBezierGridDB:
		delete(db.spiralbeziergridDBs, v.ID)
	case *SpiralCircleDB:
		delete(db.spiralcircleDBs, v.ID)
	case *SpiralCircleGridDB:
		delete(db.spiralcirclegridDBs, v.ID)
	case *SpiralLineDB:
		delete(db.spirallineDBs, v.ID)
	case *SpiralLineGridDB:
		delete(db.spirallinegridDBs, v.ID)
	case *SpiralOriginDB:
		delete(db.spiraloriginDBs, v.ID)
	case *SpiralRhombusDB:
		delete(db.spiralrhombusDBs, v.ID)
	case *SpiralRhombusGridDB:
		delete(db.spiralrhombusgridDBs, v.ID)
	case *VerticalAxisDB:
		delete(db.verticalaxisDBs, v.ID)
	default:
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AxisDB:
		db.axisDBs[v.ID] = v
		return db, nil
	case *AxisGridDB:
		db.axisgridDBs[v.ID] = v
		return db, nil
	case *BezierDB:
		db.bezierDBs[v.ID] = v
		return db, nil
	case *BezierGridDB:
		db.beziergridDBs[v.ID] = v
		return db, nil
	case *BezierGridStackDB:
		db.beziergridstackDBs[v.ID] = v
		return db, nil
	case *CircleDB:
		db.circleDBs[v.ID] = v
		return db, nil
	case *CircleGridDB:
		db.circlegridDBs[v.ID] = v
		return db, nil
	case *CursorDB:
		db.cursorDBs[v.ID] = v
		return db, nil
	case *FrontCurveDB:
		db.frontcurveDBs[v.ID] = v
		return db, nil
	case *FrontCurveStackDB:
		db.frontcurvestackDBs[v.ID] = v
		return db, nil
	case *HorizontalAxisDB:
		db.horizontalaxisDBs[v.ID] = v
		return db, nil
	case *KeyDB:
		db.keyDBs[v.ID] = v
		return db, nil
	case *ParameterDB:
		db.parameterDBs[v.ID] = v
		return db, nil
	case *RhombusDB:
		db.rhombusDBs[v.ID] = v
		return db, nil
	case *RhombusGridDB:
		db.rhombusgridDBs[v.ID] = v
		return db, nil
	case *ShapeCategoryDB:
		db.shapecategoryDBs[v.ID] = v
		return db, nil
	case *SpiralBezierDB:
		db.spiralbezierDBs[v.ID] = v
		return db, nil
	case *SpiralBezierGridDB:
		db.spiralbeziergridDBs[v.ID] = v
		return db, nil
	case *SpiralCircleDB:
		db.spiralcircleDBs[v.ID] = v
		return db, nil
	case *SpiralCircleGridDB:
		db.spiralcirclegridDBs[v.ID] = v
		return db, nil
	case *SpiralLineDB:
		db.spirallineDBs[v.ID] = v
		return db, nil
	case *SpiralLineGridDB:
		db.spirallinegridDBs[v.ID] = v
		return db, nil
	case *SpiralOriginDB:
		db.spiraloriginDBs[v.ID] = v
		return db, nil
	case *SpiralRhombusDB:
		db.spiralrhombusDBs[v.ID] = v
		return db, nil
	case *SpiralRhombusGridDB:
		db.spiralrhombusgridDBs[v.ID] = v
		return db, nil
	case *VerticalAxisDB:
		db.verticalaxisDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *AxisDB:
		if existing, ok := db.axisDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Axis github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *AxisGridDB:
		if existing, ok := db.axisgridDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AxisGrid github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *BezierDB:
		if existing, ok := db.bezierDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Bezier github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *BezierGridDB:
		if existing, ok := db.beziergridDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db BezierGrid github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *BezierGridStackDB:
		if existing, ok := db.beziergridstackDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db BezierGridStack github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *CircleDB:
		if existing, ok := db.circleDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Circle github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *CircleGridDB:
		if existing, ok := db.circlegridDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db CircleGrid github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *CursorDB:
		if existing, ok := db.cursorDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Cursor github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *FrontCurveDB:
		if existing, ok := db.frontcurveDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FrontCurve github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *FrontCurveStackDB:
		if existing, ok := db.frontcurvestackDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db FrontCurveStack github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *HorizontalAxisDB:
		if existing, ok := db.horizontalaxisDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db HorizontalAxis github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *KeyDB:
		if existing, ok := db.keyDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Key github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *ParameterDB:
		if existing, ok := db.parameterDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Parameter github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *RhombusDB:
		if existing, ok := db.rhombusDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Rhombus github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *RhombusGridDB:
		if existing, ok := db.rhombusgridDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RhombusGrid github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *ShapeCategoryDB:
		if existing, ok := db.shapecategoryDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ShapeCategory github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *SpiralBezierDB:
		if existing, ok := db.spiralbezierDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SpiralBezier github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *SpiralBezierGridDB:
		if existing, ok := db.spiralbeziergridDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SpiralBezierGrid github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *SpiralCircleDB:
		if existing, ok := db.spiralcircleDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SpiralCircle github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *SpiralCircleGridDB:
		if existing, ok := db.spiralcirclegridDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SpiralCircleGrid github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *SpiralLineDB:
		if existing, ok := db.spirallineDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SpiralLine github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *SpiralLineGridDB:
		if existing, ok := db.spirallinegridDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SpiralLineGrid github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *SpiralOriginDB:
		if existing, ok := db.spiraloriginDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SpiralOrigin github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *SpiralRhombusDB:
		if existing, ok := db.spiralrhombusDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SpiralRhombus github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *SpiralRhombusGridDB:
		if existing, ok := db.spiralrhombusgridDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SpiralRhombusGrid github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	case *VerticalAxisDB:
		if existing, ok := db.verticalaxisDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db VerticalAxis github.com/thomaspeugeot/phylotaxymusic/go, record not found")
		}
	default:
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]AxisDB:
		*ptr = make([]AxisDB, 0, len(db.axisDBs))
		for _, v := range db.axisDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]AxisGridDB:
		*ptr = make([]AxisGridDB, 0, len(db.axisgridDBs))
		for _, v := range db.axisgridDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]BezierDB:
		*ptr = make([]BezierDB, 0, len(db.bezierDBs))
		for _, v := range db.bezierDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]BezierGridDB:
		*ptr = make([]BezierGridDB, 0, len(db.beziergridDBs))
		for _, v := range db.beziergridDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]BezierGridStackDB:
		*ptr = make([]BezierGridStackDB, 0, len(db.beziergridstackDBs))
		for _, v := range db.beziergridstackDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CircleDB:
		*ptr = make([]CircleDB, 0, len(db.circleDBs))
		for _, v := range db.circleDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CircleGridDB:
		*ptr = make([]CircleGridDB, 0, len(db.circlegridDBs))
		for _, v := range db.circlegridDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]CursorDB:
		*ptr = make([]CursorDB, 0, len(db.cursorDBs))
		for _, v := range db.cursorDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FrontCurveDB:
		*ptr = make([]FrontCurveDB, 0, len(db.frontcurveDBs))
		for _, v := range db.frontcurveDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]FrontCurveStackDB:
		*ptr = make([]FrontCurveStackDB, 0, len(db.frontcurvestackDBs))
		for _, v := range db.frontcurvestackDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]HorizontalAxisDB:
		*ptr = make([]HorizontalAxisDB, 0, len(db.horizontalaxisDBs))
		for _, v := range db.horizontalaxisDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]KeyDB:
		*ptr = make([]KeyDB, 0, len(db.keyDBs))
		for _, v := range db.keyDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ParameterDB:
		*ptr = make([]ParameterDB, 0, len(db.parameterDBs))
		for _, v := range db.parameterDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RhombusDB:
		*ptr = make([]RhombusDB, 0, len(db.rhombusDBs))
		for _, v := range db.rhombusDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RhombusGridDB:
		*ptr = make([]RhombusGridDB, 0, len(db.rhombusgridDBs))
		for _, v := range db.rhombusgridDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ShapeCategoryDB:
		*ptr = make([]ShapeCategoryDB, 0, len(db.shapecategoryDBs))
		for _, v := range db.shapecategoryDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SpiralBezierDB:
		*ptr = make([]SpiralBezierDB, 0, len(db.spiralbezierDBs))
		for _, v := range db.spiralbezierDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SpiralBezierGridDB:
		*ptr = make([]SpiralBezierGridDB, 0, len(db.spiralbeziergridDBs))
		for _, v := range db.spiralbeziergridDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SpiralCircleDB:
		*ptr = make([]SpiralCircleDB, 0, len(db.spiralcircleDBs))
		for _, v := range db.spiralcircleDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SpiralCircleGridDB:
		*ptr = make([]SpiralCircleGridDB, 0, len(db.spiralcirclegridDBs))
		for _, v := range db.spiralcirclegridDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SpiralLineDB:
		*ptr = make([]SpiralLineDB, 0, len(db.spirallineDBs))
		for _, v := range db.spirallineDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SpiralLineGridDB:
		*ptr = make([]SpiralLineGridDB, 0, len(db.spirallinegridDBs))
		for _, v := range db.spirallinegridDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SpiralOriginDB:
		*ptr = make([]SpiralOriginDB, 0, len(db.spiraloriginDBs))
		for _, v := range db.spiraloriginDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SpiralRhombusDB:
		*ptr = make([]SpiralRhombusDB, 0, len(db.spiralrhombusDBs))
		for _, v := range db.spiralrhombusDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SpiralRhombusGridDB:
		*ptr = make([]SpiralRhombusGridDB, 0, len(db.spiralrhombusgridDBs))
		for _, v := range db.spiralrhombusgridDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]VerticalAxisDB:
		*ptr = make([]VerticalAxisDB, 0, len(db.verticalaxisDBs))
		for _, v := range db.verticalaxisDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *AxisDB:
		tmp, ok := db.axisDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Axis Unkown entry %d", i))
		}

		axisDB, _ := instanceDB.(*AxisDB)
		*axisDB = *tmp
		
	case *AxisGridDB:
		tmp, ok := db.axisgridDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First AxisGrid Unkown entry %d", i))
		}

		axisgridDB, _ := instanceDB.(*AxisGridDB)
		*axisgridDB = *tmp
		
	case *BezierDB:
		tmp, ok := db.bezierDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Bezier Unkown entry %d", i))
		}

		bezierDB, _ := instanceDB.(*BezierDB)
		*bezierDB = *tmp
		
	case *BezierGridDB:
		tmp, ok := db.beziergridDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First BezierGrid Unkown entry %d", i))
		}

		beziergridDB, _ := instanceDB.(*BezierGridDB)
		*beziergridDB = *tmp
		
	case *BezierGridStackDB:
		tmp, ok := db.beziergridstackDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First BezierGridStack Unkown entry %d", i))
		}

		beziergridstackDB, _ := instanceDB.(*BezierGridStackDB)
		*beziergridstackDB = *tmp
		
	case *CircleDB:
		tmp, ok := db.circleDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Circle Unkown entry %d", i))
		}

		circleDB, _ := instanceDB.(*CircleDB)
		*circleDB = *tmp
		
	case *CircleGridDB:
		tmp, ok := db.circlegridDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First CircleGrid Unkown entry %d", i))
		}

		circlegridDB, _ := instanceDB.(*CircleGridDB)
		*circlegridDB = *tmp
		
	case *CursorDB:
		tmp, ok := db.cursorDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Cursor Unkown entry %d", i))
		}

		cursorDB, _ := instanceDB.(*CursorDB)
		*cursorDB = *tmp
		
	case *FrontCurveDB:
		tmp, ok := db.frontcurveDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FrontCurve Unkown entry %d", i))
		}

		frontcurveDB, _ := instanceDB.(*FrontCurveDB)
		*frontcurveDB = *tmp
		
	case *FrontCurveStackDB:
		tmp, ok := db.frontcurvestackDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First FrontCurveStack Unkown entry %d", i))
		}

		frontcurvestackDB, _ := instanceDB.(*FrontCurveStackDB)
		*frontcurvestackDB = *tmp
		
	case *HorizontalAxisDB:
		tmp, ok := db.horizontalaxisDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First HorizontalAxis Unkown entry %d", i))
		}

		horizontalaxisDB, _ := instanceDB.(*HorizontalAxisDB)
		*horizontalaxisDB = *tmp
		
	case *KeyDB:
		tmp, ok := db.keyDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Key Unkown entry %d", i))
		}

		keyDB, _ := instanceDB.(*KeyDB)
		*keyDB = *tmp
		
	case *ParameterDB:
		tmp, ok := db.parameterDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Parameter Unkown entry %d", i))
		}

		parameterDB, _ := instanceDB.(*ParameterDB)
		*parameterDB = *tmp
		
	case *RhombusDB:
		tmp, ok := db.rhombusDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Rhombus Unkown entry %d", i))
		}

		rhombusDB, _ := instanceDB.(*RhombusDB)
		*rhombusDB = *tmp
		
	case *RhombusGridDB:
		tmp, ok := db.rhombusgridDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RhombusGrid Unkown entry %d", i))
		}

		rhombusgridDB, _ := instanceDB.(*RhombusGridDB)
		*rhombusgridDB = *tmp
		
	case *ShapeCategoryDB:
		tmp, ok := db.shapecategoryDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ShapeCategory Unkown entry %d", i))
		}

		shapecategoryDB, _ := instanceDB.(*ShapeCategoryDB)
		*shapecategoryDB = *tmp
		
	case *SpiralBezierDB:
		tmp, ok := db.spiralbezierDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SpiralBezier Unkown entry %d", i))
		}

		spiralbezierDB, _ := instanceDB.(*SpiralBezierDB)
		*spiralbezierDB = *tmp
		
	case *SpiralBezierGridDB:
		tmp, ok := db.spiralbeziergridDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SpiralBezierGrid Unkown entry %d", i))
		}

		spiralbeziergridDB, _ := instanceDB.(*SpiralBezierGridDB)
		*spiralbeziergridDB = *tmp
		
	case *SpiralCircleDB:
		tmp, ok := db.spiralcircleDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SpiralCircle Unkown entry %d", i))
		}

		spiralcircleDB, _ := instanceDB.(*SpiralCircleDB)
		*spiralcircleDB = *tmp
		
	case *SpiralCircleGridDB:
		tmp, ok := db.spiralcirclegridDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SpiralCircleGrid Unkown entry %d", i))
		}

		spiralcirclegridDB, _ := instanceDB.(*SpiralCircleGridDB)
		*spiralcirclegridDB = *tmp
		
	case *SpiralLineDB:
		tmp, ok := db.spirallineDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SpiralLine Unkown entry %d", i))
		}

		spirallineDB, _ := instanceDB.(*SpiralLineDB)
		*spirallineDB = *tmp
		
	case *SpiralLineGridDB:
		tmp, ok := db.spirallinegridDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SpiralLineGrid Unkown entry %d", i))
		}

		spirallinegridDB, _ := instanceDB.(*SpiralLineGridDB)
		*spirallinegridDB = *tmp
		
	case *SpiralOriginDB:
		tmp, ok := db.spiraloriginDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SpiralOrigin Unkown entry %d", i))
		}

		spiraloriginDB, _ := instanceDB.(*SpiralOriginDB)
		*spiraloriginDB = *tmp
		
	case *SpiralRhombusDB:
		tmp, ok := db.spiralrhombusDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SpiralRhombus Unkown entry %d", i))
		}

		spiralrhombusDB, _ := instanceDB.(*SpiralRhombusDB)
		*spiralrhombusDB = *tmp
		
	case *SpiralRhombusGridDB:
		tmp, ok := db.spiralrhombusgridDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SpiralRhombusGrid Unkown entry %d", i))
		}

		spiralrhombusgridDB, _ := instanceDB.(*SpiralRhombusGridDB)
		*spiralrhombusgridDB = *tmp
		
	case *VerticalAxisDB:
		tmp, ok := db.verticalaxisDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First VerticalAxis Unkown entry %d", i))
		}

		verticalaxisDB, _ := instanceDB.(*VerticalAxisDB)
		*verticalaxisDB = *tmp
		
	default:
		return nil, errors.New("github.com/thomaspeugeot/phylotaxymusic/go, Unkown type")
	}
	
	return db, nil
}

