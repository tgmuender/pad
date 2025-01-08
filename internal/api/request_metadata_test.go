package api

//func TestRequestMetadata_extractAuthentication(t *testing.T) {
//	type fields struct {
//		headerValues []string
//	}
//
//	metadata := RequestMetadata{headerValues: make([]string, 0)}
//
//	tests := []struct {
//		name   string
//		fields fields
//		want   *AuthenticatedIdentity
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &RequestMetadata{
//				headerValues: tt.fields.headerValues,
//			}
//			if got := m.extractAuthentication(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("extractAuthentication() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
