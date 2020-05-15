
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

type TBoundLabel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBoundLabel(owner IComponent) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = BoundLabel_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(b, (*TBoundLabel).Free)
    return b
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsBoundLabel(obj interface{}) *TBoundLabel {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TBoundLabel{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsBoundLabel.
func BoundLabelFromInst(inst uintptr) *TBoundLabel {
    return AsBoundLabel(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsBoundLabel.
func BoundLabelFromObj(obj IObject) *TBoundLabel {
    return AsBoundLabel(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsBoundLabel.
func BoundLabelFromUnsafePointer(ptr unsafe.Pointer) *TBoundLabel {
    return AsBoundLabel(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (b *TBoundLabel) Free() {
    if b.instance != 0 {
        BoundLabel_Free(b.instance)
        b.instance, b.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBoundLabel) Instance() uintptr {
    return b.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBoundLabel) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBoundLabel) IsValid() bool {
    return b.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (b *TBoundLabel) Is() TIs {
    return TIs(b.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (b *TBoundLabel) As() TAs {
//    return TAs(b.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBoundLabelClass() TClass {
    return BoundLabel_StaticClassType()
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (b *TBoundLabel) BringToFront() {
    BoundLabel_BringToFront(b.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (b *TBoundLabel) ClientToScreen(Point TPoint) TPoint {
    return BoundLabel_ClientToScreen(b.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (b *TBoundLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return BoundLabel_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (b *TBoundLabel) Dragging() bool {
    return BoundLabel_Dragging(b.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (b *TBoundLabel) HasParent() bool {
    return BoundLabel_HasParent(b.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (b *TBoundLabel) Hide() {
    BoundLabel_Hide(b.instance)
}

// CN: 要求重绘。
// EN: Redraw.
func (b *TBoundLabel) Invalidate() {
    BoundLabel_Invalidate(b.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (b *TBoundLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return BoundLabel_Perform(b.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (b *TBoundLabel) Refresh() {
    BoundLabel_Refresh(b.instance)
}

// CN: 重绘。
// EN: Repaint.
func (b *TBoundLabel) Repaint() {
    BoundLabel_Repaint(b.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (b *TBoundLabel) ScreenToClient(Point TPoint) TPoint {
    return BoundLabel_ScreenToClient(b.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (b *TBoundLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return BoundLabel_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (b *TBoundLabel) SendToBack() {
    BoundLabel_SendToBack(b.instance)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (b *TBoundLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    BoundLabel_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 显示控件。
// EN: Show control.
func (b *TBoundLabel) Show() {
    BoundLabel_Show(b.instance)
}

// CN: 控件更新。
// EN: Update.
func (b *TBoundLabel) Update() {
    BoundLabel_Update(b.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (b *TBoundLabel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return BoundLabel_GetTextBuf(b.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (b *TBoundLabel) GetTextLen() int32 {
    return BoundLabel_GetTextLen(b.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (b *TBoundLabel) SetTextBuf(Buffer string) {
    BoundLabel_SetTextBuf(b.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (b *TBoundLabel) FindComponent(AName string) *TComponent {
    return AsComponent(BoundLabel_FindComponent(b.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TBoundLabel) GetNamePath() string {
    return BoundLabel_GetNamePath(b.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TBoundLabel) Assign(Source IObject) {
    BoundLabel_Assign(b.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBoundLabel) ClassType() TClass {
    return BoundLabel_ClassType(b.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBoundLabel) ClassName() string {
    return BoundLabel_ClassName(b.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBoundLabel) InstanceSize() int32 {
    return BoundLabel_InstanceSize(b.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBoundLabel) InheritsFrom(AClass TClass) bool {
    return BoundLabel_InheritsFrom(b.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBoundLabel) Equals(Obj IObject) bool {
    return BoundLabel_Equals(b.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBoundLabel) GetHashCode() int32 {
    return BoundLabel_GetHashCode(b.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (b *TBoundLabel) ToString() string {
    return BoundLabel_ToString(b.instance)
}

func (b *TBoundLabel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    BoundLabel_AnchorToNeighbour(b.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (b *TBoundLabel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    BoundLabel_AnchorParallel(b.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (b *TBoundLabel) AnchorHorizontalCenterTo(ASibling IControl) {
    BoundLabel_AnchorHorizontalCenterTo(b.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (b *TBoundLabel) AnchorVerticalCenterTo(ASibling IControl) {
    BoundLabel_AnchorVerticalCenterTo(b.instance, CheckPtr(ASibling))
}

func (b *TBoundLabel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    BoundLabel_AnchorAsAlign(b.instance, ATheAlign , ASpace)
}

func (b *TBoundLabel) AnchorClient(ASpace int32) {
    BoundLabel_AnchorClient(b.instance, ASpace)
}

func (b *TBoundLabel) BiDiMode() TBiDiMode {
    return BoundLabel_GetBiDiMode(b.instance)
}

func (b *TBoundLabel) SetBiDiMode(value TBiDiMode) {
    BoundLabel_SetBiDiMode(b.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (b *TBoundLabel) Caption() string {
    return BoundLabel_GetCaption(b.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (b *TBoundLabel) SetCaption(value string) {
    BoundLabel_SetCaption(b.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (b *TBoundLabel) Color() TColor {
    return BoundLabel_GetColor(b.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (b *TBoundLabel) SetColor(value TColor) {
    BoundLabel_SetColor(b.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (b *TBoundLabel) DragCursor() TCursor {
    return BoundLabel_GetDragCursor(b.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (b *TBoundLabel) SetDragCursor(value TCursor) {
    BoundLabel_SetDragCursor(b.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (b *TBoundLabel) DragMode() TDragMode {
    return BoundLabel_GetDragMode(b.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (b *TBoundLabel) SetDragMode(value TDragMode) {
    BoundLabel_SetDragMode(b.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (b *TBoundLabel) Font() *TFont {
    return AsFont(BoundLabel_GetFont(b.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (b *TBoundLabel) SetFont(value *TFont) {
    BoundLabel_SetFont(b.instance, CheckPtr(value))
}

// CN: 获取高度。
// EN: Get height.
func (b *TBoundLabel) Height() int32 {
    return BoundLabel_GetHeight(b.instance)
}

// CN: 设置高度。
// EN: Set height.
func (b *TBoundLabel) SetHeight(value int32) {
    BoundLabel_SetHeight(b.instance, value)
}

// CN: 获取左边位置。
// EN: Get Left position.
func (b *TBoundLabel) Left() int32 {
    return BoundLabel_GetLeft(b.instance)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (b *TBoundLabel) ParentColor() bool {
    return BoundLabel_GetParentColor(b.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (b *TBoundLabel) SetParentColor(value bool) {
    BoundLabel_SetParentColor(b.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (b *TBoundLabel) ParentFont() bool {
    return BoundLabel_GetParentFont(b.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (b *TBoundLabel) SetParentFont(value bool) {
    BoundLabel_SetParentFont(b.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (b *TBoundLabel) ParentShowHint() bool {
    return BoundLabel_GetParentShowHint(b.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (b *TBoundLabel) SetParentShowHint(value bool) {
    BoundLabel_SetParentShowHint(b.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (b *TBoundLabel) PopupMenu() *TPopupMenu {
    return AsPopupMenu(BoundLabel_GetPopupMenu(b.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (b *TBoundLabel) SetPopupMenu(value IComponent) {
    BoundLabel_SetPopupMenu(b.instance, CheckPtr(value))
}

func (b *TBoundLabel) ShowAccelChar() bool {
    return BoundLabel_GetShowAccelChar(b.instance)
}

func (b *TBoundLabel) SetShowAccelChar(value bool) {
    BoundLabel_SetShowAccelChar(b.instance, value)
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (b *TBoundLabel) ShowHint() bool {
    return BoundLabel_GetShowHint(b.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (b *TBoundLabel) SetShowHint(value bool) {
    BoundLabel_SetShowHint(b.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (b *TBoundLabel) Top() int32 {
    return BoundLabel_GetTop(b.instance)
}

func (b *TBoundLabel) Layout() TTextLayout {
    return BoundLabel_GetLayout(b.instance)
}

func (b *TBoundLabel) SetLayout(value TTextLayout) {
    BoundLabel_SetLayout(b.instance, value)
}

// CN: 获取自动换行。
// EN: Get Automatic line break.
func (b *TBoundLabel) WordWrap() bool {
    return BoundLabel_GetWordWrap(b.instance)
}

// CN: 设置自动换行。
// EN: Set Automatic line break.
func (b *TBoundLabel) SetWordWrap(value bool) {
    BoundLabel_SetWordWrap(b.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (b *TBoundLabel) Width() int32 {
    return BoundLabel_GetWidth(b.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (b *TBoundLabel) SetWidth(value int32) {
    BoundLabel_SetWidth(b.instance, value)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (b *TBoundLabel) SetOnClick(fn TNotifyEvent) {
    BoundLabel_SetOnClick(b.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (b *TBoundLabel) SetOnDblClick(fn TNotifyEvent) {
    BoundLabel_SetOnDblClick(b.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (b *TBoundLabel) SetOnDragDrop(fn TDragDropEvent) {
    BoundLabel_SetOnDragDrop(b.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (b *TBoundLabel) SetOnDragOver(fn TDragOverEvent) {
    BoundLabel_SetOnDragOver(b.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (b *TBoundLabel) SetOnEndDrag(fn TEndDragEvent) {
    BoundLabel_SetOnEndDrag(b.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (b *TBoundLabel) SetOnMouseDown(fn TMouseEvent) {
    BoundLabel_SetOnMouseDown(b.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (b *TBoundLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    BoundLabel_SetOnMouseMove(b.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (b *TBoundLabel) SetOnMouseUp(fn TMouseEvent) {
    BoundLabel_SetOnMouseUp(b.instance, fn)
}

// CN: 获取画布。
// EN: .
func (b *TBoundLabel) Canvas() *TCanvas {
    return AsCanvas(BoundLabel_GetCanvas(b.instance))
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TBoundLabel) Enabled() bool {
    return BoundLabel_GetEnabled(b.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TBoundLabel) SetEnabled(value bool) {
    BoundLabel_SetEnabled(b.instance, value)
}

func (b *TBoundLabel) Action() *TAction {
    return AsAction(BoundLabel_GetAction(b.instance))
}

func (b *TBoundLabel) SetAction(value IComponent) {
    BoundLabel_SetAction(b.instance, CheckPtr(value))
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (b *TBoundLabel) Align() TAlign {
    return BoundLabel_GetAlign(b.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (b *TBoundLabel) SetAlign(value TAlign) {
    BoundLabel_SetAlign(b.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (b *TBoundLabel) Anchors() TAnchors {
    return BoundLabel_GetAnchors(b.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (b *TBoundLabel) SetAnchors(value TAnchors) {
    BoundLabel_SetAnchors(b.instance, value)
}

func (b *TBoundLabel) BoundsRect() TRect {
    return BoundLabel_GetBoundsRect(b.instance)
}

func (b *TBoundLabel) SetBoundsRect(value TRect) {
    BoundLabel_SetBoundsRect(b.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (b *TBoundLabel) ClientHeight() int32 {
    return BoundLabel_GetClientHeight(b.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (b *TBoundLabel) SetClientHeight(value int32) {
    BoundLabel_SetClientHeight(b.instance, value)
}

func (b *TBoundLabel) ClientOrigin() TPoint {
    return BoundLabel_GetClientOrigin(b.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (b *TBoundLabel) ClientRect() TRect {
    return BoundLabel_GetClientRect(b.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (b *TBoundLabel) ClientWidth() int32 {
    return BoundLabel_GetClientWidth(b.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (b *TBoundLabel) SetClientWidth(value int32) {
    BoundLabel_SetClientWidth(b.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (b *TBoundLabel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(BoundLabel_GetConstraints(b.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (b *TBoundLabel) SetConstraints(value *TSizeConstraints) {
    BoundLabel_SetConstraints(b.instance, CheckPtr(value))
}

// CN: 获取控件状态。
// EN: Get control state.
func (b *TBoundLabel) ControlState() TControlState {
    return BoundLabel_GetControlState(b.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (b *TBoundLabel) SetControlState(value TControlState) {
    BoundLabel_SetControlState(b.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (b *TBoundLabel) ControlStyle() TControlStyle {
    return BoundLabel_GetControlStyle(b.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (b *TBoundLabel) SetControlStyle(value TControlStyle) {
    BoundLabel_SetControlStyle(b.instance, value)
}

func (b *TBoundLabel) Floating() bool {
    return BoundLabel_GetFloating(b.instance)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TBoundLabel) Visible() bool {
    return BoundLabel_GetVisible(b.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TBoundLabel) SetVisible(value bool) {
    BoundLabel_SetVisible(b.instance, value)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TBoundLabel) Parent() *TWinControl {
    return AsWinControl(BoundLabel_GetParent(b.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TBoundLabel) SetParent(value IWinControl) {
    BoundLabel_SetParent(b.instance, CheckPtr(value))
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TBoundLabel) Cursor() TCursor {
    return BoundLabel_GetCursor(b.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TBoundLabel) SetCursor(value TCursor) {
    BoundLabel_SetCursor(b.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (b *TBoundLabel) Hint() string {
    return BoundLabel_GetHint(b.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (b *TBoundLabel) SetHint(value string) {
    BoundLabel_SetHint(b.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBoundLabel) ComponentCount() int32 {
    return BoundLabel_GetComponentCount(b.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (b *TBoundLabel) ComponentIndex() int32 {
    return BoundLabel_GetComponentIndex(b.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (b *TBoundLabel) SetComponentIndex(value int32) {
    BoundLabel_SetComponentIndex(b.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBoundLabel) Owner() *TComponent {
    return AsComponent(BoundLabel_GetOwner(b.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBoundLabel) Name() string {
    return BoundLabel_GetName(b.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBoundLabel) SetName(value string) {
    BoundLabel_SetName(b.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBoundLabel) Tag() int {
    return BoundLabel_GetTag(b.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBoundLabel) SetTag(value int) {
    BoundLabel_SetTag(b.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (b *TBoundLabel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(BoundLabel_GetAnchorSideLeft(b.instance))
}

// CN: 设置左边锚点。
// EN: .
func (b *TBoundLabel) SetAnchorSideLeft(value *TAnchorSide) {
    BoundLabel_SetAnchorSideLeft(b.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (b *TBoundLabel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(BoundLabel_GetAnchorSideTop(b.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (b *TBoundLabel) SetAnchorSideTop(value *TAnchorSide) {
    BoundLabel_SetAnchorSideTop(b.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (b *TBoundLabel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(BoundLabel_GetAnchorSideRight(b.instance))
}

// CN: 设置右边锚点。
// EN: .
func (b *TBoundLabel) SetAnchorSideRight(value *TAnchorSide) {
    BoundLabel_SetAnchorSideRight(b.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (b *TBoundLabel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(BoundLabel_GetAnchorSideBottom(b.instance))
}

// CN: 设置底边锚点。
// EN: .
func (b *TBoundLabel) SetAnchorSideBottom(value *TAnchorSide) {
    BoundLabel_SetAnchorSideBottom(b.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (b *TBoundLabel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(BoundLabel_GetBorderSpacing(b.instance))
}

// CN: 设置边框间距。
// EN: .
func (b *TBoundLabel) SetBorderSpacing(value *TControlBorderSpacing) {
    BoundLabel_SetBorderSpacing(b.instance, CheckPtr(value))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBoundLabel) Components(AIndex int32) *TComponent {
    return AsComponent(BoundLabel_GetComponents(b.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (b *TBoundLabel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(BoundLabel_GetAnchorSide(b.instance, AKind))
}

