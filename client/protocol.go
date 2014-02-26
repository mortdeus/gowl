package client

type message struct {
	name      string
	signature string
	types     []meta
}

type meta struct {
	name    string
	version int
	methods []message
	events  []message
}

var displayMeta = meta{
	name: "wl_display", version: 1,
	methods: []message{
		{"sync", "n", []meta{callbackMeta}},
		{"get_registry", "n", []meta{registryMeta}},
	},
	events: []message{
		{"error", "ous", nil},
		{"delete_id", "u", nil}},
}
var registryMeta = meta{
	name: "wl_registry", version: 1,
	methods: []message{
		{"bind", "usun", nil}},
	events: []message{
		{"global", "usu", nil},
		{"global_remove", "u", nil}},
}
var callbackMeta = meta{
	name: "wl_callback", version: 1,
	methods: nil,
	events: []message{
		{"done", "u", nil}},
}
var shmPoolMeta = meta{
	name: "wl_shm_pool", version: 1,
	methods: []message{
		{"create_buffer", "niiiiu", []meta{bufferMeta, nil, nil, nil, nil, nil}},
		{"destroy", "", nil},
		{"resize", "i", nil}},
	events: nil,
}
var shmMeta = meta{
	name: "wl_shm", version: 1,
	methods: []message{
		{"create_pool", "nhi", []meta{shmPoolMeta, nil, nil}}},
	events: []message{
		{"format", "u", nil}},
}

var bufferMeta = meta{
	name: "wl_buffer", version: 1,
	methods: []message{
		{"destroy", "", nil}},
	events: []message{
		{"release", "", nil}},
}
var outputMeta = meta{
	name: "wl_output", version: 2,
	methods: nil,
	events: []message{
		{"geometry", "iiiiissi", nil},
		{"mode", "uiii", nil},
		{"done", "2", nil},
		{"scale", "2i", nil}},
}

var dataOfferMeta = meta{
	name: "wl_data_offer", version: 1,
	methods: []message{
		{"accept", "u?s", nil},
		{"receive", "sh", nil},
		{"destroy", "", nil}},
	events: []message{
		{"offer", "s", nil}},
}

var dataSourceMeta = meta{
	name: "wl_data_source", version: 1,
	methods: []message{
		{"offer", "s", nil},
		{"destroy", "", nil}},
	events: []message{
		{"target", "?s", nil},
		{"send", "sh", nil},
		{"cancelled", "", nil}},
}

var dataDeviceMeta = meta{
	name: "wl_data_device", version: 1,
	methods: []message{
		{"start_drag", "?oo?ou", []meta{dataSourceMeta, surfaceMeta, surfaceMeta, nil}},
		{"set_selection", "?ou", []meta{dataSourceMeta, nil}}},
	events: []message{
		{"data_offer", "n", []meta{dataOfferMeta}},
		{"enter", "uoff?o", []meta{nil, surfaceMeta, nil, nil, dataOfferMeta}},
		{"leave", "", nil},
		{"motion", "uff", nil},
		{"drop", "", nil},
		{"selection", "?o", []meta{dataOfferMeta}},
	},
}
var dataDeviceManagerMeta = meta{
	name: "wl_data_device_manager", version: 1,
	methods: []message{
		{"create_data_source", "n", []meta{dataOfferMeta}},
		{"get_data_device", "no", []meta{dataDeviceMeta, seatMeta}}},
	events: nil,
}
var shellMeta = meta{
	name: "wl_shell", version: 1,
	methods: []message{
		{"get_shell_surface", "no", []meta{shellSurfaceMeta, surfaceMeta}}},
	events: nil,
}

var shellSurfaceMeta = meta{
	name: "wl_shell_surface", version: 1,
	methods: []message{
		{"pong", "u", nil},
		{"move", "ou", []meta{seatMeta, nil}},
		{"resize", "ouu", []meta{seatMeta, nil, nil}},
		{"set_toplevel", "", nil},
		{"set_transient", "oiiu", []meta{surfaceMeta, nil, nil, nil}},
		{"set_fullscreen", "uu?o", []meta{nil, nil, outputMeta}},
		{"set_popup", "ouoiiu", []meta{seatMeta, nil, surfaceMeta, nil, nil, nil}},
		{"set_maximized", "?o", []meta{outputMeta}},
		{"set_title", "s", nil},
		{"set_class", "s", nil},
	},
	events: []message{
		{"ping", "u", nil},
		{"configure", "uii", nil},
		{"popup_done", "", nil}},
}
var compositorMeta = meta{
	name: "wl_compositor", version: 3,
	methods: []message{
		{"create_surface", "n", []meta{surfaceMeta}},
		{"create_region", "n", []meta{regionMeta}}},
	events: nil,
}

var regionMeta = meta{
	name: "wl_region", version: 1,
	methods: []message{
		{"destroy", "", nil},
		{"add", "iiii", nil},
		{"subtract", "iiii", nil}},
	events: nil,
}

var surfaceMeta = meta{
	name: "wl_surface", version: 3,
	methods: []message{
		{"destroy", "", nil},
		{"attach", "?oii", []meta{bufferMeta, nil, nil}},
		{"damage", "iiii", nil},
		{"frame", "n", []meta{callbackMeta}},
		{"set_opaque_region", "?o", []meta{regionMeta}},
		{"set_input_region", "?o", []meta{regionMeta}},
		{"commit", "", nil},
		{"set_buffer_transform", "2i", nil},
		{"set_buffer_scale", "3i", nil}},
	events: []message{
		{"enter", "o", []meta{outputMeta}},
		{"leave", "o", []meta{outputMeta}}},
}
var seatMeta = meta{
	name: "wl_seat", version: 3,
	methods: []message{
		{"get_pointer", "n", []meta{pointerMeta}},
		{"get_keyboard", "n", []meta{keyboardMeta}},
		{"get_touch", "n", []meta{touchMeta}}},
	events: []message{
		{"capabilities", "u", nil},
		{"name", "2s", nil}},
}
var pointerMeta = meta{
	name: "wl_pointer", version: 3,
	methods: []message{
		{"set_cursor", "u?oii", []meta{nil, surfaceMeta, nil, nil}},
		{"release", "3", nil}},
	events: []message{
		{"enter", "3uoff", []meta{nil, surfaceMeta, nil, nil}},
		{"leave", "3uo", []meta{nil, surfaceMeta}},
		{"motion", "3uff", nil},
		{"button", "3uuuu", nil},
		{"axis", "3uuf", nil}},
}
var keyboardMeta = meta{
	name: "wl_keyboard", version: 3,
	methods: []message{
		{"release", "3", nil}},
	events: []message{
		{"keymap", "3uhu", nil},
		{"enter", "3uoa", []meta{nil, surfaceMeta, nil}},
		{"leave", "3uo", []meta{nil, surfaceMeta}},
		{"key", "3uuuu", nil},
		{"modifiers", "3uuuuu", nil}},
}
var touchMeta = meta{
	name: "wl_touch", version: 3,
	methods: []message{
		{"release", "3", nil}},
	events: []message{
		{"down", "3uuoiff", []meta{nil, nil, surfaceMeta, nil, nil, nil}},
		{"up", "3uui", nil},
		{"motion", "3uiff", nil},
		{"frame", "3", nil},
		{"cancel", "3", nil}},
}
var subcompositorMeta = meta{
	name: "wl_subcompositor", version: 1,
	methods: []message{
		{"destroy", "", nil},
		{"get_subsurface", "noo", []meta{subsurfaceMeta, surfaceMeta, surfaceMeta}}},
	events: nil,
}
var subsurfaceMeta = meta{
	name: "wl_subsurface", version: 1,
	methods: []message{
		{"destroy", "", nil},
		{"set_position", "ii", nil},
		{"place_above", "o", []meta{surfaceMeta}},
		{"place_below", "o", []meta{surfaceMeta}},
		{"set_sync", "", nil},
		{"set_desync", "", nil}},
	events: nil,
}
