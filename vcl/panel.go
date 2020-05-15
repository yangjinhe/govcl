
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

type TPanel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPanel(owner IComponent) *TPanel {
    p := new(TPanel)
    p.instance = Panel_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(p, (*TPanel).Free)
    return p
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsPanel(obj interface{}) *TPanel {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TPanel{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsPanel.
func PanelFromInst(inst uintptr) *TPanel {
    return AsPanel(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsPanel.
func PanelFromObj(obj IObject) *TPanel {
    return AsPanel(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPanel.
func PanelFromUnsafePointer(ptr unsafe.Pointer) *TPanel {
    return AsPanel(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (p *TPanel) Free() {
    if p.instance != 0 {
        Panel_Free(p.instance)
        p.instance, p.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPanel) Instance() uintptr {
    return p.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPanel) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPanel) IsValid() bool {
    return p.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (p *TPanel) Is() TIs {
    return TIs(p.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (p *TPanel) As() TAs {
//    return TAs(p.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPanelClass() TClass {
    return Panel_StaticClassType()
}

// CN: 是否可以获得焦点。
// EN: .
func (p *TPanel) CanFocus() bool {
    return Panel_CanFocus(p.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (p *TPanel) ContainsControl(Control IControl) bool {
    return Panel_ContainsControl(p.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (p *TPanel) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(Panel_ControlAtPos(p.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (p *TPanel) DisableAlign() {
    Panel_DisableAlign(p.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (p *TPanel) EnableAlign() {
    Panel_EnableAlign(p.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (p *TPanel) FindChildControl(ControlName string) *TControl {
    return AsControl(Panel_FindChildControl(p.instance, ControlName))
}

func (p *TPanel) FlipChildren(AllLevels bool) {
    Panel_FlipChildren(p.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (p *TPanel) Focused() bool {
    return Panel_Focused(p.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (p *TPanel) HandleAllocated() bool {
    return Panel_HandleAllocated(p.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (p *TPanel) InsertControl(AControl IControl) {
    Panel_InsertControl(p.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (p *TPanel) Invalidate() {
    Panel_Invalidate(p.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (p *TPanel) RemoveControl(AControl IControl) {
    Panel_RemoveControl(p.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (p *TPanel) Realign() {
    Panel_Realign(p.instance)
}

// CN: 重绘。
// EN: Repaint.
func (p *TPanel) Repaint() {
    Panel_Repaint(p.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (p *TPanel) ScaleBy(M int32, D int32) {
    Panel_ScaleBy(p.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (p *TPanel) ScrollBy(DeltaX int32, DeltaY int32) {
    Panel_ScrollBy(p.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (p *TPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Panel_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (p *TPanel) SetFocus() {
    Panel_SetFocus(p.instance)
}

// CN: 控件更新。
// EN: Update.
func (p *TPanel) Update() {
    Panel_Update(p.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (p *TPanel) BringToFront() {
    Panel_BringToFront(p.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (p *TPanel) ClientToScreen(Point TPoint) TPoint {
    return Panel_ClientToScreen(p.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (p *TPanel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Panel_ClientToParent(p.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (p *TPanel) Dragging() bool {
    return Panel_Dragging(p.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TPanel) HasParent() bool {
    return Panel_HasParent(p.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (p *TPanel) Hide() {
    Panel_Hide(p.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (p *TPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Panel_Perform(p.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (p *TPanel) Refresh() {
    Panel_Refresh(p.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (p *TPanel) ScreenToClient(Point TPoint) TPoint {
    return Panel_ScreenToClient(p.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (p *TPanel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Panel_ParentToClient(p.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (p *TPanel) SendToBack() {
    Panel_SendToBack(p.instance)
}

// CN: 显示控件。
// EN: Show control.
func (p *TPanel) Show() {
    Panel_Show(p.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (p *TPanel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Panel_GetTextBuf(p.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (p *TPanel) GetTextLen() int32 {
    return Panel_GetTextLen(p.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (p *TPanel) SetTextBuf(Buffer string) {
    Panel_SetTextBuf(p.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TPanel) FindComponent(AName string) *TComponent {
    return AsComponent(Panel_FindComponent(p.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPanel) GetNamePath() string {
    return Panel_GetNamePath(p.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPanel) Assign(Source IObject) {
    Panel_Assign(p.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPanel) ClassType() TClass {
    return Panel_ClassType(p.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPanel) ClassName() string {
    return Panel_ClassName(p.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPanel) InstanceSize() int32 {
    return Panel_InstanceSize(p.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPanel) InheritsFrom(AClass TClass) bool {
    return Panel_InheritsFrom(p.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPanel) Equals(Obj IObject) bool {
    return Panel_Equals(p.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPanel) GetHashCode() int32 {
    return Panel_GetHashCode(p.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (p *TPanel) ToString() string {
    return Panel_ToString(p.instance)
}

func (p *TPanel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Panel_AnchorToNeighbour(p.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (p *TPanel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Panel_AnchorParallel(p.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (p *TPanel) AnchorHorizontalCenterTo(ASibling IControl) {
    Panel_AnchorHorizontalCenterTo(p.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (p *TPanel) AnchorVerticalCenterTo(ASibling IControl) {
    Panel_AnchorVerticalCenterTo(p.instance, CheckPtr(ASibling))
}

func (p *TPanel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Panel_AnchorAsAlign(p.instance, ATheAlign , ASpace)
}

func (p *TPanel) AnchorClient(ASpace int32) {
    Panel_AnchorClient(p.instance, ASpace)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (p *TPanel) Align() TAlign {
    return Panel_GetAlign(p.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (p *TPanel) SetAlign(value TAlign) {
    Panel_SetAlign(p.instance, value)
}

// CN: 获取文字对齐。
// EN: Get Text alignment.
func (p *TPanel) Alignment() TAlignment {
    return Panel_GetAlignment(p.instance)
}

// CN: 设置文字对齐。
// EN: Set Text alignment.
func (p *TPanel) SetAlignment(value TAlignment) {
    Panel_SetAlignment(p.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (p *TPanel) Anchors() TAnchors {
    return Panel_GetAnchors(p.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (p *TPanel) SetAnchors(value TAnchors) {
    Panel_SetAnchors(p.instance, value)
}

// CN: 获取自动调整大小。
// EN: .
func (p *TPanel) AutoSize() bool {
    return Panel_GetAutoSize(p.instance)
}

// CN: 设置自动调整大小。
// EN: .
func (p *TPanel) SetAutoSize(value bool) {
    Panel_SetAutoSize(p.instance, value)
}

func (p *TPanel) BevelInner() TBevelCut {
    return Panel_GetBevelInner(p.instance)
}

func (p *TPanel) SetBevelInner(value TBevelCut) {
    Panel_SetBevelInner(p.instance, value)
}

func (p *TPanel) BevelOuter() TBevelCut {
    return Panel_GetBevelOuter(p.instance)
}

func (p *TPanel) SetBevelOuter(value TBevelCut) {
    Panel_SetBevelOuter(p.instance, value)
}

func (p *TPanel) BiDiMode() TBiDiMode {
    return Panel_GetBiDiMode(p.instance)
}

func (p *TPanel) SetBiDiMode(value TBiDiMode) {
    Panel_SetBiDiMode(p.instance, value)
}

// CN: 获取边框的宽度。
// EN: .
func (p *TPanel) BorderWidth() int32 {
    return Panel_GetBorderWidth(p.instance)
}

// CN: 设置边框的宽度。
// EN: .
func (p *TPanel) SetBorderWidth(value int32) {
    Panel_SetBorderWidth(p.instance, value)
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (p *TPanel) BorderStyle() TBorderStyle {
    return Panel_GetBorderStyle(p.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (p *TPanel) SetBorderStyle(value TBorderStyle) {
    Panel_SetBorderStyle(p.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (p *TPanel) Caption() string {
    return Panel_GetCaption(p.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (p *TPanel) SetCaption(value string) {
    Panel_SetCaption(p.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (p *TPanel) Color() TColor {
    return Panel_GetColor(p.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (p *TPanel) SetColor(value TColor) {
    Panel_SetColor(p.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (p *TPanel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Panel_GetConstraints(p.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (p *TPanel) SetConstraints(value *TSizeConstraints) {
    Panel_SetConstraints(p.instance, CheckPtr(value))
}

// CN: 获取使用停靠管理。
// EN: .
func (p *TPanel) UseDockManager() bool {
    return Panel_GetUseDockManager(p.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (p *TPanel) SetUseDockManager(value bool) {
    Panel_SetUseDockManager(p.instance, value)
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (p *TPanel) DockSite() bool {
    return Panel_GetDockSite(p.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (p *TPanel) SetDockSite(value bool) {
    Panel_SetDockSite(p.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (p *TPanel) DoubleBuffered() bool {
    return Panel_GetDoubleBuffered(p.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (p *TPanel) SetDoubleBuffered(value bool) {
    Panel_SetDoubleBuffered(p.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (p *TPanel) DragCursor() TCursor {
    return Panel_GetDragCursor(p.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (p *TPanel) SetDragCursor(value TCursor) {
    Panel_SetDragCursor(p.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (p *TPanel) DragKind() TDragKind {
    return Panel_GetDragKind(p.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (p *TPanel) SetDragKind(value TDragKind) {
    Panel_SetDragKind(p.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (p *TPanel) DragMode() TDragMode {
    return Panel_GetDragMode(p.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (p *TPanel) SetDragMode(value TDragMode) {
    Panel_SetDragMode(p.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (p *TPanel) Enabled() bool {
    return Panel_GetEnabled(p.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (p *TPanel) SetEnabled(value bool) {
    Panel_SetEnabled(p.instance, value)
}

func (p *TPanel) FullRepaint() bool {
    return Panel_GetFullRepaint(p.instance)
}

func (p *TPanel) SetFullRepaint(value bool) {
    Panel_SetFullRepaint(p.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (p *TPanel) Font() *TFont {
    return AsFont(Panel_GetFont(p.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (p *TPanel) SetFont(value *TFont) {
    Panel_SetFont(p.instance, CheckPtr(value))
}

func (p *TPanel) ParentBackground() bool {
    return Panel_GetParentBackground(p.instance)
}

func (p *TPanel) SetParentBackground(value bool) {
    Panel_SetParentBackground(p.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (p *TPanel) ParentColor() bool {
    return Panel_GetParentColor(p.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (p *TPanel) SetParentColor(value bool) {
    Panel_SetParentColor(p.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (p *TPanel) ParentDoubleBuffered() bool {
    return Panel_GetParentDoubleBuffered(p.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (p *TPanel) SetParentDoubleBuffered(value bool) {
    Panel_SetParentDoubleBuffered(p.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (p *TPanel) ParentFont() bool {
    return Panel_GetParentFont(p.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (p *TPanel) SetParentFont(value bool) {
    Panel_SetParentFont(p.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (p *TPanel) ParentShowHint() bool {
    return Panel_GetParentShowHint(p.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (p *TPanel) SetParentShowHint(value bool) {
    Panel_SetParentShowHint(p.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (p *TPanel) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Panel_GetPopupMenu(p.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (p *TPanel) SetPopupMenu(value IComponent) {
    Panel_SetPopupMenu(p.instance, CheckPtr(value))
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (p *TPanel) ShowHint() bool {
    return Panel_GetShowHint(p.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (p *TPanel) SetShowHint(value bool) {
    Panel_SetShowHint(p.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (p *TPanel) TabOrder() TTabOrder {
    return Panel_GetTabOrder(p.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (p *TPanel) SetTabOrder(value TTabOrder) {
    Panel_SetTabOrder(p.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (p *TPanel) TabStop() bool {
    return Panel_GetTabStop(p.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (p *TPanel) SetTabStop(value bool) {
    Panel_SetTabStop(p.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (p *TPanel) Visible() bool {
    return Panel_GetVisible(p.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (p *TPanel) SetVisible(value bool) {
    Panel_SetVisible(p.instance, value)
}

func (p *TPanel) SetOnAlignPosition(fn TAlignPositionEvent) {
    Panel_SetOnAlignPosition(p.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (p *TPanel) SetOnClick(fn TNotifyEvent) {
    Panel_SetOnClick(p.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (p *TPanel) SetOnContextPopup(fn TContextPopupEvent) {
    Panel_SetOnContextPopup(p.instance, fn)
}

func (p *TPanel) SetOnDockDrop(fn TDockDropEvent) {
    Panel_SetOnDockDrop(p.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (p *TPanel) SetOnDblClick(fn TNotifyEvent) {
    Panel_SetOnDblClick(p.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (p *TPanel) SetOnDragDrop(fn TDragDropEvent) {
    Panel_SetOnDragDrop(p.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (p *TPanel) SetOnDragOver(fn TDragOverEvent) {
    Panel_SetOnDragOver(p.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (p *TPanel) SetOnEndDock(fn TEndDragEvent) {
    Panel_SetOnEndDock(p.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (p *TPanel) SetOnEndDrag(fn TEndDragEvent) {
    Panel_SetOnEndDrag(p.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (p *TPanel) SetOnEnter(fn TNotifyEvent) {
    Panel_SetOnEnter(p.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (p *TPanel) SetOnExit(fn TNotifyEvent) {
    Panel_SetOnExit(p.instance, fn)
}

func (p *TPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    Panel_SetOnGetSiteInfo(p.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (p *TPanel) SetOnMouseDown(fn TMouseEvent) {
    Panel_SetOnMouseDown(p.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (p *TPanel) SetOnMouseEnter(fn TNotifyEvent) {
    Panel_SetOnMouseEnter(p.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (p *TPanel) SetOnMouseLeave(fn TNotifyEvent) {
    Panel_SetOnMouseLeave(p.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (p *TPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    Panel_SetOnMouseMove(p.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (p *TPanel) SetOnMouseUp(fn TMouseEvent) {
    Panel_SetOnMouseUp(p.instance, fn)
}

// CN: 设置大小被改变事件。
// EN: .
func (p *TPanel) SetOnResize(fn TNotifyEvent) {
    Panel_SetOnResize(p.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (p *TPanel) SetOnStartDock(fn TStartDockEvent) {
    Panel_SetOnStartDock(p.instance, fn)
}

func (p *TPanel) SetOnUnDock(fn TUnDockEvent) {
    Panel_SetOnUnDock(p.instance, fn)
}

// CN: 获取依靠客户端总数。
// EN: .
func (p *TPanel) DockClientCount() int32 {
    return Panel_GetDockClientCount(p.instance)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (p *TPanel) MouseInClient() bool {
    return Panel_GetMouseInClient(p.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (p *TPanel) VisibleDockClientCount() int32 {
    return Panel_GetVisibleDockClientCount(p.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (p *TPanel) Brush() *TBrush {
    return AsBrush(Panel_GetBrush(p.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (p *TPanel) ControlCount() int32 {
    return Panel_GetControlCount(p.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TPanel) Handle() HWND {
    return Panel_GetHandle(p.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (p *TPanel) ParentWindow() HWND {
    return Panel_GetParentWindow(p.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (p *TPanel) SetParentWindow(value HWND) {
    Panel_SetParentWindow(p.instance, value)
}

func (p *TPanel) Action() *TAction {
    return AsAction(Panel_GetAction(p.instance))
}

func (p *TPanel) SetAction(value IComponent) {
    Panel_SetAction(p.instance, CheckPtr(value))
}

func (p *TPanel) BoundsRect() TRect {
    return Panel_GetBoundsRect(p.instance)
}

func (p *TPanel) SetBoundsRect(value TRect) {
    Panel_SetBoundsRect(p.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (p *TPanel) ClientHeight() int32 {
    return Panel_GetClientHeight(p.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (p *TPanel) SetClientHeight(value int32) {
    Panel_SetClientHeight(p.instance, value)
}

func (p *TPanel) ClientOrigin() TPoint {
    return Panel_GetClientOrigin(p.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (p *TPanel) ClientRect() TRect {
    return Panel_GetClientRect(p.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (p *TPanel) ClientWidth() int32 {
    return Panel_GetClientWidth(p.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (p *TPanel) SetClientWidth(value int32) {
    Panel_SetClientWidth(p.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (p *TPanel) ControlState() TControlState {
    return Panel_GetControlState(p.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (p *TPanel) SetControlState(value TControlState) {
    Panel_SetControlState(p.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (p *TPanel) ControlStyle() TControlStyle {
    return Panel_GetControlStyle(p.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (p *TPanel) SetControlStyle(value TControlStyle) {
    Panel_SetControlStyle(p.instance, value)
}

func (p *TPanel) Floating() bool {
    return Panel_GetFloating(p.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (p *TPanel) Parent() *TWinControl {
    return AsWinControl(Panel_GetParent(p.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (p *TPanel) SetParent(value IWinControl) {
    Panel_SetParent(p.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (p *TPanel) Left() int32 {
    return Panel_GetLeft(p.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (p *TPanel) SetLeft(value int32) {
    Panel_SetLeft(p.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (p *TPanel) Top() int32 {
    return Panel_GetTop(p.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (p *TPanel) SetTop(value int32) {
    Panel_SetTop(p.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (p *TPanel) Width() int32 {
    return Panel_GetWidth(p.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (p *TPanel) SetWidth(value int32) {
    Panel_SetWidth(p.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (p *TPanel) Height() int32 {
    return Panel_GetHeight(p.instance)
}

// CN: 设置高度。
// EN: Set height.
func (p *TPanel) SetHeight(value int32) {
    Panel_SetHeight(p.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (p *TPanel) Cursor() TCursor {
    return Panel_GetCursor(p.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (p *TPanel) SetCursor(value TCursor) {
    Panel_SetCursor(p.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (p *TPanel) Hint() string {
    return Panel_GetHint(p.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (p *TPanel) SetHint(value string) {
    Panel_SetHint(p.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPanel) ComponentCount() int32 {
    return Panel_GetComponentCount(p.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (p *TPanel) ComponentIndex() int32 {
    return Panel_GetComponentIndex(p.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (p *TPanel) SetComponentIndex(value int32) {
    Panel_SetComponentIndex(p.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPanel) Owner() *TComponent {
    return AsComponent(Panel_GetOwner(p.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPanel) Name() string {
    return Panel_GetName(p.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPanel) SetName(value string) {
    Panel_SetName(p.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPanel) Tag() int {
    return Panel_GetTag(p.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPanel) SetTag(value int) {
    Panel_SetTag(p.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (p *TPanel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSideLeft(p.instance))
}

// CN: 设置左边锚点。
// EN: .
func (p *TPanel) SetAnchorSideLeft(value *TAnchorSide) {
    Panel_SetAnchorSideLeft(p.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (p *TPanel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSideTop(p.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (p *TPanel) SetAnchorSideTop(value *TAnchorSide) {
    Panel_SetAnchorSideTop(p.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (p *TPanel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSideRight(p.instance))
}

// CN: 设置右边锚点。
// EN: .
func (p *TPanel) SetAnchorSideRight(value *TAnchorSide) {
    Panel_SetAnchorSideRight(p.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (p *TPanel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSideBottom(p.instance))
}

// CN: 设置底边锚点。
// EN: .
func (p *TPanel) SetAnchorSideBottom(value *TAnchorSide) {
    Panel_SetAnchorSideBottom(p.instance, CheckPtr(value))
}

func (p *TPanel) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(Panel_GetChildSizing(p.instance))
}

func (p *TPanel) SetChildSizing(value *TControlChildSizing) {
    Panel_SetChildSizing(p.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (p *TPanel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Panel_GetBorderSpacing(p.instance))
}

// CN: 设置边框间距。
// EN: .
func (p *TPanel) SetBorderSpacing(value *TControlBorderSpacing) {
    Panel_SetBorderSpacing(p.instance, CheckPtr(value))
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (p *TPanel) DockClients(Index int32) *TControl {
    return AsControl(Panel_GetDockClients(p.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (p *TPanel) Controls(Index int32) *TControl {
    return AsControl(Panel_GetControls(p.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPanel) Components(AIndex int32) *TComponent {
    return AsComponent(Panel_GetComponents(p.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (p *TPanel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Panel_GetAnchorSide(p.instance, AKind))
}

