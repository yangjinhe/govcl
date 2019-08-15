procedure SizeConstraints_Assign(AObj: TSizeConstraints; Source: TPersistent); extdecl;
begin
  AObj.Assign(Source);
end;

function SizeConstraints_GetNamePath(AObj: TSizeConstraints): PChar; extdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function SizeConstraints_ClassType(AObj: TSizeConstraints): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function SizeConstraints_ClassName(AObj: TSizeConstraints): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function SizeConstraints_InstanceSize(AObj: TSizeConstraints): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function SizeConstraints_InheritsFrom(AObj: TSizeConstraints; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function SizeConstraints_Equals(AObj: TSizeConstraints; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function SizeConstraints_GetHashCode(AObj: TSizeConstraints): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function SizeConstraints_ToString(AObj: TSizeConstraints): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

procedure SizeConstraints_SetOnChange(AObj: TSizeConstraints; AEventId: NativeUInt); extdecl;
begin
  {$DEFINE EventName := OnChange}
  {$DEFINE EventMethodName := OnChange}
  EventMethodCode
end;

function SizeConstraints_GetMaxHeight(AObj: TSizeConstraints): TConstraintSize; extdecl;
begin
  Result :=  AObj.MaxHeight;
end;

procedure SizeConstraints_SetMaxHeight(AObj: TSizeConstraints; AValue: TConstraintSize); extdecl;
begin
  AObj.MaxHeight := AValue;
end;

function SizeConstraints_GetMaxWidth(AObj: TSizeConstraints): TConstraintSize; extdecl;
begin
  Result :=  AObj.MaxWidth;
end;

procedure SizeConstraints_SetMaxWidth(AObj: TSizeConstraints; AValue: TConstraintSize); extdecl;
begin
  AObj.MaxWidth := AValue;
end;

function SizeConstraints_GetMinHeight(AObj: TSizeConstraints): TConstraintSize; extdecl;
begin
  Result :=  AObj.MinHeight;
end;

procedure SizeConstraints_SetMinHeight(AObj: TSizeConstraints; AValue: TConstraintSize); extdecl;
begin
  AObj.MinHeight := AValue;
end;

function SizeConstraints_GetMinWidth(AObj: TSizeConstraints): TConstraintSize; extdecl;
begin
  Result :=  AObj.MinWidth;
end;

procedure SizeConstraints_SetMinWidth(AObj: TSizeConstraints; AValue: TConstraintSize); extdecl;
begin
  AObj.MinWidth := AValue;
end;

function SizeConstraints_StaticClassType: TClass; extdecl;
begin
  Result := TSizeConstraints;
end;

exports
  SizeConstraints_Assign,
  SizeConstraints_GetNamePath,
  SizeConstraints_ClassType,
  SizeConstraints_ClassName,
  SizeConstraints_InstanceSize,
  SizeConstraints_InheritsFrom,
  SizeConstraints_Equals,
  SizeConstraints_GetHashCode,
  SizeConstraints_ToString,
  SizeConstraints_SetOnChange,
  SizeConstraints_GetMaxHeight,
  SizeConstraints_SetMaxHeight,
  SizeConstraints_GetMaxWidth,
  SizeConstraints_SetMaxWidth,
  SizeConstraints_GetMinHeight,
  SizeConstraints_SetMinHeight,
  SizeConstraints_GetMinWidth,
  SizeConstraints_SetMinWidth,
  SizeConstraints_StaticClassType;
