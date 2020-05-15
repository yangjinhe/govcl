
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TLabel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewLabel(owner IComponent) *TLabel {
    l := new(TLabel)
    l.instance = Label_Create(CheckPtr(owner))
    l.ptr = unsafe.Pointer(l.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(l, (*TLabel).Free)
    return l
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsLabel(obj interface{}) *TLabel {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TLabel{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsLabel.
func LabelFromInst(inst uintptr) *TLabel {
    return AsLabel(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsLabel.
func LabelFromObj(obj IObject) *TLabel {
    return AsLabel(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsLabel.
func LabelFromUnsafePointer(ptr unsafe.Pointer) *TLabel {
    return AsLabel(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (l *TLabel) Free() {
    if l.instance != 0 {
        Label_Free(l.instance)
        l.instance, l.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TLabel) Instance() uintptr {
    return l.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TLabel) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TLabel) IsValid() bool {
    return l.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (l *TLabel) Is() TIs {
    return TIs(l.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (l *TLabel) As() TAs {
//    return TAs(l.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TLabelClass() TClass {
    return Label_StaticClassType()
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (l *TLabel) BringToFront() {
    Label_BringToFront(l.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (l *TLabel) ClientToScreen(Point TPoint) TPoint {
    return Label_ClientToScreen(l.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (l *TLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Label_ClientToParent(l.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (l *TLabel) Dragging() bool {
    return Label_Dragging(l.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (l *TLabel) HasParent() bool {
    return Label_HasParent(l.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (l *TLabel) Hide() {
    Label_Hide(l.instance)
}

// CN: 要求重绘。
// EN: Redraw.
func (l *TLabel) Invalidate() {
    Label_Invalidate(l.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (l *TLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Label_Perform(l.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (l *TLabel) Refresh() {
    Label_Refresh(l.instance)
}

// CN: 重绘。
// EN: Repaint.
func (l *TLabel) Repaint() {
    Label_Repaint(l.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (l *TLabel) ScreenToClient(Point TPoint) TPoint {
    return Label_ScreenToClient(l.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (l *TLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Label_ParentToClient(l.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (l *TLabel) SendToBack() {
    Label_SendToBack(l.instance)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (l *TLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Label_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 显示控件。
// EN: Show control.
func (l *TLabel) Show() {
    Label_Show(l.instance)
}

// CN: 控件更新。
// EN: Update.
func (l *TLabel) Update() {
    Label_Update(l.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (l *TLabel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Label_GetTextBuf(l.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (l *TLabel) GetTextLen() int32 {
    return Label_GetTextLen(l.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (l *TLabel) SetTextBuf(Buffer string) {
    Label_SetTextBuf(l.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (l *TLabel) FindComponent(AName string) *TComponent {
    return AsComponent(Label_FindComponent(l.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TLabel) GetNamePath() string {
    return Label_GetNamePath(l.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TLabel) Assign(Source IObject) {
    Label_Assign(l.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TLabel) ClassType() TClass {
    return Label_ClassType(l.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TLabel) ClassName() string {
    return Label_ClassName(l.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TLabel) InstanceSize() int32 {
    return Label_InstanceSize(l.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TLabel) InheritsFrom(AClass TClass) bool {
    return Label_InheritsFrom(l.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TLabel) Equals(Obj IObject) bool {
    return Label_Equals(l.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TLabel) GetHashCode() int32 {
    return Label_GetHashCode(l.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (l *TLabel) ToString() string {
    return Label_ToString(l.instance)
}

func (l *TLabel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Label_AnchorToNeighbour(l.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (l *TLabel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Label_AnchorParallel(l.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (l *TLabel) AnchorHorizontalCenterTo(ASibling IControl) {
    Label_AnchorHorizontalCenterTo(l.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (l *TLabel) AnchorVerticalCenterTo(ASibling IControl) {
    Label_AnchorVerticalCenterTo(l.instance, CheckPtr(ASibling))
}

func (l *TLabel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Label_AnchorAsAlign(l.instance, ATheAlign , ASpace)
}

func (l *TLabel) AnchorClient(ASpace int32) {
    Label_AnchorClient(l.instance, ASpace)
}

func (l *TLabel) OptimalFill() bool {
    return Label_GetOptimalFill(l.instance)
}

func (l *TLabel) SetOptimalFill(value bool) {
    Label_SetOptimalFill(l.instance, value)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (l *TLabel) Align() TAlign {
    return Label_GetAlign(l.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (l *TLabel) SetAlign(value TAlign) {
    Label_SetAlign(l.instance, value)
}

// CN: 获取文字对齐。
// EN: Get Text alignment.
func (l *TLabel) Alignment() TAlignment {
    return Label_GetAlignment(l.instance)
}

// CN: 设置文字对齐。
// EN: Set Text alignment.
func (l *TLabel) SetAlignment(value TAlignment) {
    Label_SetAlignment(l.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (l *TLabel) Anchors() TAnchors {
    return Label_GetAnchors(l.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (l *TLabel) SetAnchors(value TAnchors) {
    Label_SetAnchors(l.instance, value)
}

// CN: 获取自动调整大小。
// EN: .
func (l *TLabel) AutoSize() bool {
    return Label_GetAutoSize(l.instance)
}

// CN: 设置自动调整大小。
// EN: .
func (l *TLabel) SetAutoSize(value bool) {
    Label_SetAutoSize(l.instance, value)
}

func (l *TLabel) BiDiMode() TBiDiMode {
    return Label_GetBiDiMode(l.instance)
}

func (l *TLabel) SetBiDiMode(value TBiDiMode) {
    Label_SetBiDiMode(l.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (l *TLabel) Caption() string {
    return Label_GetCaption(l.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (l *TLabel) SetCaption(value string) {
    Label_SetCaption(l.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (l *TLabel) Color() TColor {
    return Label_GetColor(l.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (l *TLabel) SetColor(value TColor) {
    Label_SetColor(l.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (l *TLabel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Label_GetConstraints(l.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (l *TLabel) SetConstraints(value *TSizeConstraints) {
    Label_SetConstraints(l.instance, CheckPtr(value))
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (l *TLabel) DragCursor() TCursor {
    return Label_GetDragCursor(l.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (l *TLabel) SetDragCursor(value TCursor) {
    Label_SetDragCursor(l.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (l *TLabel) DragKind() TDragKind {
    return Label_GetDragKind(l.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (l *TLabel) SetDragKind(value TDragKind) {
    Label_SetDragKind(l.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (l *TLabel) DragMode() TDragMode {
    return Label_GetDragMode(l.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (l *TLabel) SetDragMode(value TDragMode) {
    Label_SetDragMode(l.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (l *TLabel) Enabled() bool {
    return Label_GetEnabled(l.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (l *TLabel) SetEnabled(value bool) {
    Label_SetEnabled(l.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (l *TLabel) Font() *TFont {
    return AsFont(Label_GetFont(l.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (l *TLabel) SetFont(value *TFont) {
    Label_SetFont(l.instance, CheckPtr(value))
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (l *TLabel) ParentColor() bool {
    return Label_GetParentColor(l.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (l *TLabel) SetParentColor(value bool) {
    Label_SetParentColor(l.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (l *TLabel) ParentFont() bool {
    return Label_GetParentFont(l.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (l *TLabel) SetParentFont(value bool) {
    Label_SetParentFont(l.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (l *TLabel) ParentShowHint() bool {
    return Label_GetParentShowHint(l.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (l *TLabel) SetParentShowHint(value bool) {
    Label_SetParentShowHint(l.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (l *TLabel) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Label_GetPopupMenu(l.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (l *TLabel) SetPopupMenu(value IComponent) {
    Label_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TLabel) ShowAccelChar() bool {
    return Label_GetShowAccelChar(l.instance)
}

func (l *TLabel) SetShowAccelChar(value bool) {
    Label_SetShowAccelChar(l.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (l *TLabel) ShowHint() bool {
    return Label_GetShowHint(l.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (l *TLabel) SetShowHint(value bool) {
    Label_SetShowHint(l.instance, value)
}

// CN: 获取透明。
// EN: Get transparent.
func (l *TLabel) Transparent() bool {
    return Label_GetTransparent(l.instance)
}

// CN: 设置透明。
// EN: Set transparent.
func (l *TLabel) SetTransparent(value bool) {
    Label_SetTransparent(l.instance, value)
}

func (l *TLabel) Layout() TTextLayout {
    return Label_GetLayout(l.instance)
}

func (l *TLabel) SetLayout(value TTextLayout) {
    Label_SetLayout(l.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (l *TLabel) Visible() bool {
    return Label_GetVisible(l.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (l *TLabel) SetVisible(value bool) {
    Label_SetVisible(l.instance, value)
}

// CN: 获取自动换行。
// EN: Get Automatic line break.
func (l *TLabel) WordWrap() bool {
    return Label_GetWordWrap(l.instance)
}

// CN: 设置自动换行。
// EN: Set Automatic line break.
func (l *TLabel) SetWordWrap(value bool) {
    Label_SetWordWrap(l.instance, value)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (l *TLabel) SetOnClick(fn TNotifyEvent) {
    Label_SetOnClick(l.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (l *TLabel) SetOnContextPopup(fn TContextPopupEvent) {
    Label_SetOnContextPopup(l.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (l *TLabel) SetOnDblClick(fn TNotifyEvent) {
    Label_SetOnDblClick(l.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (l *TLabel) SetOnDragDrop(fn TDragDropEvent) {
    Label_SetOnDragDrop(l.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (l *TLabel) SetOnDragOver(fn TDragOverEvent) {
    Label_SetOnDragOver(l.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (l *TLabel) SetOnEndDrag(fn TEndDragEvent) {
    Label_SetOnEndDrag(l.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (l *TLabel) SetOnMouseDown(fn TMouseEvent) {
    Label_SetOnMouseDown(l.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (l *TLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    Label_SetOnMouseMove(l.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (l *TLabel) SetOnMouseUp(fn TMouseEvent) {
    Label_SetOnMouseUp(l.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (l *TLabel) SetOnMouseEnter(fn TNotifyEvent) {
    Label_SetOnMouseEnter(l.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (l *TLabel) SetOnMouseLeave(fn TNotifyEvent) {
    Label_SetOnMouseLeave(l.instance, fn)
}

// CN: 获取画布。
// EN: .
func (l *TLabel) Canvas() *TCanvas {
    return AsCanvas(Label_GetCanvas(l.instance))
}

func (l *TLabel) Action() *TAction {
    return AsAction(Label_GetAction(l.instance))
}

func (l *TLabel) SetAction(value IComponent) {
    Label_SetAction(l.instance, CheckPtr(value))
}

func (l *TLabel) BoundsRect() TRect {
    return Label_GetBoundsRect(l.instance)
}

func (l *TLabel) SetBoundsRect(value TRect) {
    Label_SetBoundsRect(l.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (l *TLabel) ClientHeight() int32 {
    return Label_GetClientHeight(l.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (l *TLabel) SetClientHeight(value int32) {
    Label_SetClientHeight(l.instance, value)
}

func (l *TLabel) ClientOrigin() TPoint {
    return Label_GetClientOrigin(l.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (l *TLabel) ClientRect() TRect {
    return Label_GetClientRect(l.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (l *TLabel) ClientWidth() int32 {
    return Label_GetClientWidth(l.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (l *TLabel) SetClientWidth(value int32) {
    Label_SetClientWidth(l.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (l *TLabel) ControlState() TControlState {
    return Label_GetControlState(l.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (l *TLabel) SetControlState(value TControlState) {
    Label_SetControlState(l.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (l *TLabel) ControlStyle() TControlStyle {
    return Label_GetControlStyle(l.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (l *TLabel) SetControlStyle(value TControlStyle) {
    Label_SetControlStyle(l.instance, value)
}

func (l *TLabel) Floating() bool {
    return Label_GetFloating(l.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (l *TLabel) Parent() *TWinControl {
    return AsWinControl(Label_GetParent(l.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (l *TLabel) SetParent(value IWinControl) {
    Label_SetParent(l.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (l *TLabel) Left() int32 {
    return Label_GetLeft(l.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (l *TLabel) SetLeft(value int32) {
    Label_SetLeft(l.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (l *TLabel) Top() int32 {
    return Label_GetTop(l.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (l *TLabel) SetTop(value int32) {
    Label_SetTop(l.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (l *TLabel) Width() int32 {
    return Label_GetWidth(l.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (l *TLabel) SetWidth(value int32) {
    Label_SetWidth(l.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (l *TLabel) Height() int32 {
    return Label_GetHeight(l.instance)
}

// CN: 设置高度。
// EN: Set height.
func (l *TLabel) SetHeight(value int32) {
    Label_SetHeight(l.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (l *TLabel) Cursor() TCursor {
    return Label_GetCursor(l.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (l *TLabel) SetCursor(value TCursor) {
    Label_SetCursor(l.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (l *TLabel) Hint() string {
    return Label_GetHint(l.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (l *TLabel) SetHint(value string) {
    Label_SetHint(l.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (l *TLabel) ComponentCount() int32 {
    return Label_GetComponentCount(l.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (l *TLabel) ComponentIndex() int32 {
    return Label_GetComponentIndex(l.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (l *TLabel) SetComponentIndex(value int32) {
    Label_SetComponentIndex(l.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TLabel) Owner() *TComponent {
    return AsComponent(Label_GetOwner(l.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (l *TLabel) Name() string {
    return Label_GetName(l.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (l *TLabel) SetName(value string) {
    Label_SetName(l.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (l *TLabel) Tag() int {
    return Label_GetTag(l.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (l *TLabel) SetTag(value int) {
    Label_SetTag(l.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (l *TLabel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSideLeft(l.instance))
}

// CN: 设置左边锚点。
// EN: .
func (l *TLabel) SetAnchorSideLeft(value *TAnchorSide) {
    Label_SetAnchorSideLeft(l.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (l *TLabel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSideTop(l.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (l *TLabel) SetAnchorSideTop(value *TAnchorSide) {
    Label_SetAnchorSideTop(l.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (l *TLabel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSideRight(l.instance))
}

// CN: 设置右边锚点。
// EN: .
func (l *TLabel) SetAnchorSideRight(value *TAnchorSide) {
    Label_SetAnchorSideRight(l.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (l *TLabel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSideBottom(l.instance))
}

// CN: 设置底边锚点。
// EN: .
func (l *TLabel) SetAnchorSideBottom(value *TAnchorSide) {
    Label_SetAnchorSideBottom(l.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (l *TLabel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Label_GetBorderSpacing(l.instance))
}

// CN: 设置边框间距。
// EN: .
func (l *TLabel) SetBorderSpacing(value *TControlBorderSpacing) {
    Label_SetBorderSpacing(l.instance, CheckPtr(value))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (l *TLabel) Components(AIndex int32) *TComponent {
    return AsComponent(Label_GetComponents(l.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (l *TLabel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Label_GetAnchorSide(l.instance, AKind))
}

