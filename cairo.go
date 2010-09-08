package cairo

// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe";
)

// cairo_status_t values
const (
	StatusSuccess	= iota;
	StatusNoMemory;
	StatusInvalidRestore;
	StatusInvalidPopGroup;
	StatusNoCurrentPoint;
	StatusInvalidMatrix;
	StatusInvalidStatus;
	StatusNullPointer;
	StatusInvalidString;
	StatusInvalidPathData;
	StatusReadError;
	StatusWriteError;
	StatusSurfaceFinished;
	StatusSurfaceTypeMismatch;
	StatusPatternTypeMismatch;
	StatusInvalidContent;
	StatusInvalidFormat;
	StatusInvalidVisual;
	StatusFileNotFound;
	StatusInvalidDash;
	StatusInvalidDscComment;
	StatusInvalidIndex;
	StatusClipNotRepresentable;
	StatusTempFileError;
	StatusInvalidStride;
	StatusFontTypeMismatch;
	StatusUserFontImmutable;
	StatusUserFontError;
	StatusNegativeCount;
	StatusInvalidClusters;
	StatusInvalidSlant;
	StatusInvalidWeight;
)

// cairo_content_t values
const (
	ContentColor	= iota * 0x1000;
	ContentAlpha;
	ContentColorAlpha;
)

// cairo_operator_t values
const (
	OperatorClear	= iota;
	OperatorSource;
	OperatorOver;
	OperatorIn;
	OperatorOut;
	OperatorAtop;
	OperatorDest;
	OperatorDestOver;
	OperatorDestIn;
	OperatorDestOut;
	OperatorDestAtop;
	OperatorXor;
	OperatorAdd;
	OperatorSaturate;
)

// cairo_antialias_t values
const (
	AntialiasDefault	= iota;
	AntialiasNone;
	AntialiasGray;
	AntialiasSubpixel;
)

// cairo_fill_rule_t values
const (
	FillRuleWinding	= iota;
	FillRuleEvenOdd;
)

// cairo_line_cap_t values
const (
	LineCapButt	= iota;
	LineCapRound;
	LineCapSquare;
)

// cairo_line_cap_join_t values
const (
	LineJoinMiter	= iota;
	LineJoinRound;
	LineJoinBevel;
)

// cairo_text_cluster_flag_t value(s)
const (
	TextClusterFlagBackward = iota;
)

// cairo_font_slant_t values
const (
	FontSlantNormal	= iota;
	FontSlantItalic;
	FontSlantOblique;
)

// cairo_font_weight_t values
const (
	FontWeightNormal	= iota;
	FontWeightBold;
)

// cairo_subpixel_order_t values
const (
	SubpixelOrderDefault	= iota;
	SubpixelOrderRGB;
	SubpixelOrderBGR;
	SubpixelOrderVRGB;
	SubpixelOrderVBGR;
)

// cairo_hint_style_t values
const (
	HintStyleDefault	= iota;
	HintStyleNone;
	HintStyleSlight;
	HintStyleMedium;
	HintStyleFull;
)

// cairo_hint_metrics_t values
const (
	HintMetricsDefault	= iota;
	HintMetricsOff;
	HintMetricsOn;
)

// cairo_font_type_t values
const (
	FontTypeToy	= iota;
	FontTypeFT;
	FontTypeWin32;
	FontTypeQuartz;
	FontTypeUser;
)

// cairo_path_data_type_t values
const (
	PathMoveTo	= iota;
	PathLineTo;
	PathCurveTo;
	PathClosePath;
)

// cairo_surface_type_t values
const (
	SurfaceTypeImage	= iota;
	SurfaceTypePDF;
	SurfaceTypePS;
	SurfaceTypeXlib;
	SurfaceTypeXCB;
	SurfaceTypeGlitz;
	SurfaceTypeQuartz;
	SurfaceTypeWin32;
	SurfaceTypeBeOS;
	SurfaceTypeDirectFB;
	SurfaceTypeSVG;
	SurfaceTypeOS2;
	SurfaceTypeWin32Printing;
	SurfaceTypeQuartzImage;
)

// cairo_format_t values
const (
	FormatArgB32	= iota;
	FormatRGB24;
	FormatA8;
	FormatA1;
	_;	// reserved for deprecated value
)

// cairo_pattern_type_t values
const (
	PatternTypeSolid	= iota;
	PatternTypeSurface;
	PatternTypeLinear;
	PatternTypeRadial;
)

// cairo_extend_t value
const (
	ExtendNone	= iota;
	ExtendRepeat;
	ExtendReflect;
	ExtendPad;
)

// cairo_filter_t values
const (
	FilterFast	= iota;
	FilterGood;
	FilterBest;
	FilterNearest;
	FilterBilinear;
	FilterGaussian;
)

// Utility functions

func cairobool2bool(flag C.cairo_bool_t) bool {
	if int(flag) > 0 {
		return true
	}
	return false;
}


type Matrix struct {
	matrix *C.cairo_matrix_t;
}

type Pattern struct {
	pattern *C.cairo_pattern_t;
}

// Golang struct to hold both a cairo surface and a cairo context
type Surface struct {
	surface	*C.cairo_surface_t;
	context	*C.cairo_t;
}

func Version() int	{ return int(C.cairo_version()) }

func NewSurface(format, width, height int) *Surface {
	surface := new(Surface);
	surface.surface = C.cairo_image_surface_create(C.cairo_format_t(format), C.int(width), C.int(height));
	surface.context = C.cairo_create(surface.surface);
	return surface;
}

func (self *Surface) Save()	{ C.cairo_save(self.context) }

func (self *Surface) Restore()	{ C.cairo_restore(self.context) }

func (self *Surface) PushGroup()	{ C.cairo_push_group(self.context) }

func (self *Surface) PushGroupWithContent(content_t int) {
	C.cairo_push_group_with_content(self.context, C.cairo_content_t(content_t))
}

func (self *Surface) PopGroup() (pattern *Pattern) {
	pattern = new(Pattern);
	pattern.pattern = C.cairo_pop_group(self.context);
	return;
}

func (self *Surface) PopGroupToSource()	{ C.cairo_pop_group_to_source(self.context) }

func (self *Surface) SetOperator(operator_t int) {
	C.cairo_set_operator(self.context, C.cairo_operator_t(operator_t))
}

func (self *Surface) SetSource(pattern *Pattern) {
	C.cairo_set_source(self.context, pattern.pattern)
}

func (self *Surface) SetSourceRGB(red, green, blue float64) {
	C.cairo_set_source_rgb(self.context, C.double(red), C.double(green), C.double(blue))
}

func (self *Surface) SetSourceRGBA(red, green, blue, alpha float64) {
	C.cairo_set_source_rgba(self.context, C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

func (self *Surface) SetSourceSurface(surface *Surface, x, y float64) {
	C.cairo_set_source_surface(self.context, surface.surface, C.double(x), C.double(y))
}

func (self *Surface) SetTolerance(tolerance float64) {
	C.cairo_set_tolerance(self.context, C.double(tolerance))
}

func (self *Surface) SetAntialias(antialias_t int) {
	C.cairo_set_antialias(self.context, C.cairo_antialias_t(antialias_t))
}

func (self *Surface) SetFillRule(fill_rule_t int) {
	C.cairo_set_fill_rule(self.context, C.cairo_fill_rule_t(fill_rule_t))
}

func (self *Surface) SetLineWidth(width float64) {
	C.cairo_set_line_width(self.context, C.double(width))
}

func (self *Surface) SetLineCap(line_cap_t int) {
	C.cairo_set_line_cap(self.context, C.cairo_line_cap_t(line_cap_t))
}

func (self *Surface) SetLineJoin(line_join_t int) {
	C.cairo_set_line_join(self.context, C.cairo_line_join_t(line_join_t))
}

// TODO: Figure out how to convert a slice into C
/*func (self *Surface) SetDash(dashes []float64, num_dashes int, offset float64){
    dashesp := (*C.double)(&dashes);
    C.cairo_set_dash(self.context, dashesp, C.int(num_dashes), C.double(offset));
}
*/
func (self *Surface) SetMiterLimit(limit float64) {
	C.cairo_set_miter_limit(self.context, C.double(limit))
}

func (self *Surface) Translate(tx, ty float64) {
	C.cairo_translate(self.context, C.double(tx), C.double(ty))
}

func (self *Surface) Scale(sx, sy float64) {
	C.cairo_scale(self.context, C.double(sx), C.double(sy))
}

func (self *Surface) Rotate(angle float64)	{ C.cairo_rotate(self.context, C.double(angle)) }

func (self *Surface) Transform(matrix *Matrix) {
	C.cairo_transform(self.context, matrix.matrix)
}

func (self *Surface) SetMatrix(matrix *Matrix) {
	C.cairo_set_matrix(self.context, matrix.matrix)
}

func (self *Surface) IdentityMatrix()	{ C.cairo_identity_matrix(self.context) }

func (self *Surface) UserToDevice(x, y float64) (x1, y1 float64) {
	ux, uy := (*C.double)(&x), (*C.double)(&y);
	C.cairo_user_to_device(self.context, ux, uy);
	x1, y1 = float64(*ux), float64(*uy);
	return;
}

func (self *Surface) UserToDeviceDistance(dx, dy float64) (dx1, dy1 float64) {
	ux, uy := (*C.double)(&dx), (*C.double)(&dy);
	C.cairo_user_to_device_distance(self.context, ux, uy);
	dx1, dy1 = float64(*ux), float64(*uy);
	return;
}

// path creation methods

func (self *Surface) NewPath()	{ C.cairo_new_path(self.context) }

func (self *Surface) MoveTo(x, y float64) {
	C.cairo_move_to(self.context, C.double(x), C.double(y))
}

func (self *Surface) NewSubPath()	{ C.cairo_new_sub_path(self.context) }

func (self *Surface) LineTo(x, y float64) {
	C.cairo_line_to(self.context, C.double(x), C.double(y))
}

func (self *Surface) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_curve_to(self.context,
		C.double(x1), C.double(y1),
		C.double(x2), C.double(y2),
		C.double(x3), C.double(y3))
}

func (self *Surface) Arc(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc(self.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}

func (self *Surface) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc_negative(self.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}

func (self *Surface) RelMoveTo(dx, dy float64) {
	C.cairo_rel_move_to(self.context, C.double(dx), C.double(dy))
}

func (self *Surface) RelLineTo(dx, dy float64) {
	C.cairo_rel_line_to(self.context, C.double(dx), C.double(dy))
}

func (self *Surface) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
	C.cairo_rel_curve_to(self.context,
		C.double(dx1), C.double(dy1),
		C.double(dx2), C.double(dy2),
		C.double(dx3), C.double(dy3))
}

func (self *Surface) Rectangle(x, y, width, height float64) {
	C.cairo_rectangle(self.context,
		C.double(x), C.double(y),
		C.double(width), C.double(height))
}

func (self *Surface) ClosePath()	{ C.cairo_close_path(self.context) }

func (self *Surface) PathExtents() (left, top, right, bottom float64) {
	var x1, y1, x2, y2 float64;
	px1 := (*C.double)(&x1);
	py1 := (*C.double)(&y1);
	px2 := (*C.double)(&x2);
	py2 := (*C.double)(&y2);
	C.cairo_path_extents(self.context, px1, py1, px2, py2);
	left = float64(*px1);
	top = float64(*py1);
	right = float64(*px2);
	bottom = float64(*py2);
	return;
}

// Painting methods

func (self *Surface) Paint()	{ C.cairo_paint(self.context) }

func (self *Surface) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(self.context, C.double(alpha))
}

func (self *Surface) Mask(pattern *Pattern)	{ C.cairo_mask(self.context, pattern.pattern) }

func (self *Surface) MaskSurface(surface *Surface, surface_x, surface_y float64) {
	C.cairo_mask_surface(self.context, surface.surface, C.double(surface_x), C.double(surface_y))
}

func (self *Surface) Stroke()	{ C.cairo_stroke(self.context) }

func (self *Surface) StrokePreserve()	{ C.cairo_stroke_preserve(self.context) }

func (self *Surface) Fill()	{ C.cairo_fill(self.context) }

func (self *Surface) FillPreserve()	{ C.cairo_fill_preserve(self.context) }

func (self *Surface) CopyPage()	{ C.cairo_copy_page(self.context) }

func (self *Surface) ShowPage()	{ C.cairo_show_page(self.context) }

// Insideness testing

func (self *Surface) InStroke(x, y float64) bool {
	ret := C.cairo_in_stroke(self.context, C.double(x), C.double(y));
	return cairobool2bool(ret);
}

func (self *Surface) InFill(x, y float64) bool {
	ret := C.cairo_in_fill(self.context, C.double(x), C.double(y));
	return cairobool2bool(ret);
}

// Rectangular extents

func (self *Surface) StrokeExtents() (left, top, right, bottom float64) {
	var x1, y1, x2, y2 float64;
	px1 := (*C.double)(&x1);
	py1 := (*C.double)(&y1);
	px2 := (*C.double)(&x2);
	py2 := (*C.double)(&y2);
	C.cairo_stroke_extents(self.context, px1, py1, px2, py2);
	left = float64(*px1);
	top = float64(*py1);
	right = float64(*px2);
	bottom = float64(*py2);
	return;
}

func (self *Surface) FillExtents() (left, top, right, bottom float64) {
	var x1, y1, x2, y2 float64;
	px1 := (*C.double)(&x1);
	py1 := (*C.double)(&y1);
	px2 := (*C.double)(&x2);
	py2 := (*C.double)(&y2);
	C.cairo_fill_extents(self.context, px1, py1, px2, py2);
	left = float64(*px1);
	top = float64(*py1);
	right = float64(*px2);
	bottom = float64(*py2);
	return;
}

// Clipping methods

func (self *Surface) ResetClip()	{ C.cairo_reset_clip(self.context) }

func (self *Surface) Clip()	{ C.cairo_clip(self.context) }

func (self *Surface) ClipPreserve()	{ C.cairo_clip_preserve(self.context) }

func (self *Surface) ClipExtents() (left, top, right, bottom float64) {
	var x1, y1, x2, y2 float64;
	px1 := (*C.double)(&x1);
	py1 := (*C.double)(&y1);
	px2 := (*C.double)(&x2);
	py2 := (*C.double)(&y2);
	C.cairo_clip_extents(self.context, px1, py1, px2, py2);
	left = float64(*px1);
	top = float64(*py1);
	right = float64(*px2);
	bottom = float64(*py2);
	return;
}

func (self *Surface) SelectFontFace(name string, font_slant_t, font_weight_t int) {
	p := C.CString(name);
	C.cairo_select_font_face(self.context, p, C.cairo_font_slant_t(font_slant_t), C.cairo_font_weight_t(font_weight_t));
	C.free(unsafe.Pointer(p));
}

func (self *Surface) SetFontSize(size float64) {
	C.cairo_set_font_size(self.context, C.double(size))
}

func (self *Surface) ShowText(text string) {
	p := C.CString(text);
	C.cairo_show_text(self.context, p);
	C.free(unsafe.Pointer(p));
}

func (self *Surface) Finish()	{ C.cairo_destroy(self.context) }

func (self *Surface) WriteToPNG(filename string) {
	p := C.CString(filename);
	C.cairo_surface_write_to_png(self.surface, p);
	C.free(unsafe.Pointer(p));
}

func (self *Surface) destroy()	{ C.cairo_surface_destroy(self.surface) }
