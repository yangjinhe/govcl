
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

type TPrinter struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPrinter() *TPrinter {
    p := new(TPrinter)
    p.instance = Printer_Create()
    p.ptr = unsafe.Pointer(p.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(p, (*TPrinter).Free)
    return p
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsPrinter(obj interface{}) *TPrinter {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TPrinter{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsPrinter.
func PrinterFromInst(inst uintptr) *TPrinter {
    return AsPrinter(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsPrinter.
func PrinterFromObj(obj IObject) *TPrinter {
    return AsPrinter(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPrinter.
func PrinterFromUnsafePointer(ptr unsafe.Pointer) *TPrinter {
    return AsPrinter(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (p *TPrinter) Free() {
    if p.instance != 0 {
        Printer_Free(p.instance)
        p.instance, p.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPrinter) Instance() uintptr {
    return p.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPrinter) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPrinter) IsValid() bool {
    return p.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (p *TPrinter) Is() TIs {
    return TIs(p.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (p *TPrinter) As() TAs {
//    return TAs(p.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPrinterClass() TClass {
    return Printer_StaticClassType()
}

func (p *TPrinter) Abort() {
    Printer_Abort(p.instance)
}

func (p *TPrinter) BeginDoc() {
    Printer_BeginDoc(p.instance)
}

func (p *TPrinter) EndDoc() {
    Printer_EndDoc(p.instance)
}

func (p *TPrinter) NewPage() {
    Printer_NewPage(p.instance)
}

// CN: 刷新控件。
// EN: Refresh control.
func (p *TPrinter) Refresh() {
    Printer_Refresh(p.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPrinter) ClassType() TClass {
    return Printer_ClassType(p.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPrinter) ClassName() string {
    return Printer_ClassName(p.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPrinter) InstanceSize() int32 {
    return Printer_InstanceSize(p.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPrinter) InheritsFrom(AClass TClass) bool {
    return Printer_InheritsFrom(p.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPrinter) Equals(Obj IObject) bool {
    return Printer_Equals(p.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPrinter) GetHashCode() int32 {
    return Printer_GetHashCode(p.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (p *TPrinter) ToString() string {
    return Printer_ToString(p.instance)
}

func (p *TPrinter) Aborted() bool {
    return Printer_GetAborted(p.instance)
}

// CN: 获取画布。
// EN: .
func (p *TPrinter) Canvas() *TCanvas {
    return AsCanvas(Printer_GetCanvas(p.instance))
}

func (p *TPrinter) Copies() int32 {
    return Printer_GetCopies(p.instance)
}

func (p *TPrinter) SetCopies(value int32) {
    Printer_SetCopies(p.instance, value)
}

func (p *TPrinter) Fonts() *TStrings {
    return AsStrings(Printer_GetFonts(p.instance))
}

func (p *TPrinter) Orientation() TPrinterOrientation {
    return Printer_GetOrientation(p.instance)
}

func (p *TPrinter) SetOrientation(value TPrinterOrientation) {
    Printer_SetOrientation(p.instance, value)
}

func (p *TPrinter) PageHeight() int32 {
    return Printer_GetPageHeight(p.instance)
}

func (p *TPrinter) PageWidth() int32 {
    return Printer_GetPageWidth(p.instance)
}

func (p *TPrinter) PageNumber() int32 {
    return Printer_GetPageNumber(p.instance)
}

func (p *TPrinter) PrinterIndex() int32 {
    return Printer_GetPrinterIndex(p.instance)
}

func (p *TPrinter) SetPrinterIndex(value int32) {
    Printer_SetPrinterIndex(p.instance, value)
}

func (p *TPrinter) Printing() bool {
    return Printer_GetPrinting(p.instance)
}

func (p *TPrinter) Printers() *TStrings {
    return AsStrings(Printer_GetPrinters(p.instance))
}

func (p *TPrinter) Title() string {
    return Printer_GetTitle(p.instance)
}

func (p *TPrinter) SetTitle(value string) {
    Printer_SetTitle(p.instance, value)
}

