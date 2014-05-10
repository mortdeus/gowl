package client

type message struct {
	name      string
	signature string
	types     []*typ
}

type typ struct {
	name    string
	version int
	methods []message
	events  []message
}

var displayMeta = &typ{
	name: "wl_display", version: 1,
	methods: []message{
		{"sync", "n", []*typ{callbackMeta}},
		{"get_registry", "n", []*typ{registryMeta}},
	},
	events: []message{
		{"error", "ous", nil},
		{"delete_id", "u", nil}},
}
var registryMeta = &typ{
	name: "wl_registry", version: 1,
	methods: []message{
		{"bind", "usun", nil}},
	events: []message{
		{"global", "usu", nil},
		{"global_remove", "u", nil}},
}
var callbackMeta = &typ{
	name: "wl_callback", version: 1,
	methods: nil,
	events: []message{
		{"done", "u", nil}},
}
var shmPoolMeta = &typ{
	name: "wl_shm_pool", version: 1,
	methods: []message{
		{"create_buffer", "niiiiu", []*typ{bufferMeta, nil, nil, nil, nil, nil}},
		{"destroy", "", nil},
		{"resize", "i", nil}},
	events: nil,
}
var shmMeta = &typ{
	name: "wl_shm", version: 1,
	methods: []message{
		{"create_pool", "nhi", []*typ{shmPoolMeta, nil, nil}}},
	events: []message{
		{"format", "u", nil}},
}

var bufferMeta = &typ{
	name: "wl_buffer", version: 1,
	methods: []message{
		{"destroy", "", nil}},
	events: []message{
		{"release", "", nil}},
}
var outputMeta = &typ{
	name: "wl_output", version: 2,
	methods: nil,
	events: []message{
		{"geometry", "iiiiissi", nil},
		{"mode", "uiii", nil},
		{"done", "2", nil},
		{"scale", "2i", nil}},
}

var dataOfferMeta = &typ{
	name: "wl_data_offer", version: 1,
	methods: []message{
		{"accept", "u?s", nil},
		{"receive", "sh", nil},
		{"destroy", "", nil}},
	events: []message{
		{"offer", "s", nil}},
}

var dataSourceMeta = &typ{
	name: "wl_data_source", version: 1,
	methods: []message{
		{"offer", "s", nil},
		{"destroy", "", nil}},
	events: []message{
		{"target", "?s", nil},
		{"send", "sh", nil},
		{"cancelled", "", nil}},
}

var dataDeviceMeta = &typ{
	name: "wl_data_device", version: 1,
	methods: []message{
		{"start_drag", "?oo?ou", []*typ{dataSourceMeta, surfaceMeta, surfaceMeta, nil}},
		{"set_selection", "?ou", []*typ{dataSourceMeta, nil}}},
	events: []message{
		{"data_offer", "n", []*typ{dataOfferMeta}},
		{"enter", "uoff?o", []*typ{nil, surfaceMeta, nil, nil, dataOfferMeta}},
		{"leave", "", nil},
		{"motion", "uff", nil},
		{"drop", "", nil},
		{"selection", "?o", []*typ{dataOfferMeta}},
	},
}
var dataDeviceManagerMeta = &typ{
	name: "wl_data_device_manager", version: 1,
	methods: []message{
		{"create_data_source", "n", []*typ{dataOfferMeta}},
		{"get_data_device", "no", []*typ{dataDeviceMeta, seatMeta}}},
	events: nil,
}
var shellMeta = &typ{
	name: "wl_shell", version: 1,
	methods: []message{
		{"get_shell_surface", "no", []*typ{shellSurfaceMeta, surfaceMeta}}},
	events: nil,
}

var shellSurfaceMeta = &typ{
	name: "wl_shell_surface", version: 1,
	methods: []message{
		{"pong", "u", nil},
		{"move", "ou", []*typ{seatMeta, nil}},
		{"resize", "ouu", []*typ{seatMeta, nil, nil}},
		{"set_toplevel", "", nil},
		{"set_transient", "oiiu", []*typ{surfaceMeta, nil, nil, nil}},
		{"set_fullscreen", "uu?o", []*typ{nil, nil, outputMeta}},
		{"set_popup", "ouoiiu", []*typ{seatMeta, nil, surfaceMeta, nil, nil, nil}},
		{"set_maximized", "?o", []*typ{outputMeta}},
		{"set_title", "s", nil},
		{"set_class", "s", nil},
	},
	events: []message{
		{"ping", "u", nil},
		{"configure", "uii", nil},
		{"popup_done", "", nil}},
}
var compositorMeta = &typ{
	name: "wl_compositor", version: 3,
	methods: []message{
		{"create_surface", "n", []*typ{surfaceMeta}},
		{"create_region", "n", []*typ{regionMeta}}},
	events: nil,
}

var regionMeta = &typ{
	name: "wl_region", version: 1,
	methods: []message{
		{"destroy", "", nil},
		{"add", "iiii", nil},
		{"subtract", "iiii", nil}},
	events: nil,
}

var surfaceMeta = &typ{
	name: "wl_surface", version: 3,
	methods: []message{
		{"destroy", "", nil},
		{"attach", "?oii", []*typ{bufferMeta, nil, nil}},
		{"damage", "iiii", nil},
		{"frame", "n", []*typ{callbackMeta}},
		{"set_opaque_region", "?o", []*typ{regionMeta}},
		{"set_input_region", "?o", []*typ{regionMeta}},
		{"commit", "", nil},
		{"set_buffer_transform", "2i", nil},
		{"set_buffer_scale", "3i", nil}},
	events: []message{
		{"enter", "o", []*typ{outputMeta}},
		{"leave", "o", []*typ{outputMeta}}},
}
var seatMeta = &typ{
	name: "wl_seat", version: 3,
	methods: []message{
		{"get_pointer", "n", []*typ{pointerMeta}},
		{"get_keyboard", "n", []*typ{keyboardMeta}},
		{"get_touch", "n", []*typ{touchMeta}}},
	events: []message{
		{"capabilities", "u", nil},
		{"name", "2s", nil}},
}
var pointerMeta = &typ{
	name: "wl_pointer", version: 3,
	methods: []message{
		{"set_cursor", "u?oii", []*typ{nil, surfaceMeta, nil, nil}},
		{"release", "3", nil}},
	events: []message{
		{"enter", "3uoff", []*typ{nil, surfaceMeta, nil, nil}},
		{"leave", "3uo", []*typ{nil, surfaceMeta}},
		{"motion", "3uff", nil},
		{"button", "3uuuu", nil},
		{"axis", "3uuf", nil}},
}
var keyboardMeta = &typ{
	name: "wl_keyboard", version: 3,
	methods: []message{
		{"release", "3", nil}},
	events: []message{
		{"keymap", "3uhu", nil},
		{"enter", "3uoa", []*typ{nil, surfaceMeta, nil}},
		{"leave", "3uo", []*typ{nil, surfaceMeta}},
		{"key", "3uuuu", nil},
		{"modifiers", "3uuuuu", nil}},
}
var touchMeta = &typ{
	name: "wl_touch", version: 3,
	methods: []message{
		{"release", "3", nil}},
	events: []message{
		{"down", "3uuoiff", []*typ{nil, nil, surfaceMeta, nil, nil, nil}},
		{"up", "3uui", nil},
		{"motion", "3uiff", nil},
		{"frame", "3", nil},
		{"cancel", "3", nil}},
}
var subcompositorMeta = &typ{
	name: "wl_subcompositor", version: 1,
	methods: []message{
		{"destroy", "", nil},
		{"get_subsurface", "noo", []*typ{subsurfaceMeta, surfaceMeta, surfaceMeta}}},
	events: nil,
}
var subsurfaceMeta = &typ{
	name: "wl_subsurface", version: 1,
	methods: []message{
		{"destroy", "", nil},
		{"set_position", "ii", nil},
		{"place_above", "o", []*typ{surfaceMeta}},
		{"place_below", "o", []*typ{surfaceMeta}},
		{"set_sync", "", nil},
		{"set_desync", "", nil}},
	events: nil,
}
