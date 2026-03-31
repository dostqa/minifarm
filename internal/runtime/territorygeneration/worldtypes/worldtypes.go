package worldtypes

type WorldObjectType int

const (
	Grass     WorldObjectType = 0
	Err       WorldObjectType = 1
	SourceU   WorldObjectType = 11
	SourceL   WorldObjectType = 12
	SourceD   WorldObjectType = 13
	SourceR   WorldObjectType = 14
	RbedH     WorldObjectType = 15
	RbedV     WorldObjectType = 16
	RcornerUL WorldObjectType = 17
	RcornerDL WorldObjectType = 18
	RcornerDR WorldObjectType = 19
	RcornerUR WorldObjectType = 20
	OneTree   WorldObjectType = 21
	TwoTrees  WorldObjectType = 22
)
