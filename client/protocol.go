package client

var display = iface{
	name: "wl_display", version: 1,
	methods: []message{
		{"sync", "n", []*iface{&callback}},
		{"get_registry", "n", []*iface{&registry}},
	},
	events: []message{
		{"error", "ous", nil},
		{"delete_id", "u", nil}},
}
var registry = iface{
	name: "wl_registry", version: 1,
	methods: []message{
		{"bind", "usun", nil}},
	events: []message{
		{"global", "usu", nil},
		{"global_remove", "u", nil}},
}
var callback = iface{
	name: "wl_callback", version: 1,
	methods: nil,
	events: []message{
		{"done", "u", nil}},
}
var shmPool = iface{
	name: "wl_shm_pool", version: 1,
	methods: []message{
		{"create_buffer", "niiiiu", []*iface{&buffer, nil, nil, nil, nil, nil}},
		{"destroy", "", nil},
		{"resize", "i", nil}},
	events: nil,
}
var shm = iface{
	name: "wl_shm", version: 1,
	methods: []message{
		{"create_pool", "nhi", []*iface{&shmPool, nil, nil}}},
	events: []message{
		{"format", "u", nil}},
}

var buffer = iface{
	name: "wl_buffer", version: 1,
	methods: []message{
		{"destroy", "", nil}},
	events: []message{
		{"release", "", nil}},
}
var output = iface{
	name: "wl_output", version: 2,
	methods: nil,
	events: []message{
		{"geometry", "iiiiissi", nil},
		{"mode", "uiii", nil},
		{"done", "2", nil},
		{"scale", "2i", nil}},
}

var dataOffer = iface{
	name: "wl_data_offer", version: 1,
	methods: []message{
		{"accept", "u?s", nil},
		{"receive", "sh", nil},
		{"destroy", "", nil}},
	events: []message{
		{"offer", "s", nil}},
}

var dataSource = iface{
	name: "wl_data_source", version: 1,
	methods: []message{
		{"offer", "s", nil},
		{"destroy", "", nil}},
	events: []message{
		{"target", "?s", nil},
		{"send", "sh", nil},
		{"cancelled", "", nil}},
}

var dataDevice = iface{
	name: "wl_data_device", version: 1,
	methods: []message{
		{"start_drag", "?oo?ou", []*iface{&dataSource, &surface, &surface, nil}},
		{"set_selection", "?ou", []*iface{&dataSource, nil}}},
	events: []message{
		{"data_offer", "n", []*iface{&dataOffer}},
		{"enter", "uoff?o", []*iface{nil, &surface, nil, nil, &dataOffer}},
		{"leave", "", nil},
		{"motion", "uff", nil},
		{"drop", "", nil},
		{"selection", "?o", []*iface{&dataOffer}},
	},
}
var dataDeviceManager = iface{
	name: "wl_data_device_manager", version: 1,
	methods: []message{
		{"create_data_source", "n", []*iface{&dataOffer}},
		{"get_data_device", "no", []*iface{&dataDevice, &seat}}},
	events: nil,
}
var shell = iface{
	name: "wl_shell", version: 1,
	methods: []message{
		{"get_shell_surface", "no", []*iface{&shellSurface, &surface}}},
	events: nil,
}

var shellSurface = iface{
	name: "wl_shell_surface", version: 1,
	methods: []message{
		{"pong", "u", nil},
		{"move", "ou", []*iface{&seat, nil}},
		{"resize", "ouu", []*iface{&seat, nil, nil}},
		{"set_toplevel", "", nil},
		{"set_transient", "oiiu", []*iface{&surface, nil, nil, nil}},
		{"set_fullscreen", "uu?o", []*iface{nil, nil, &output}},
		{"set_popup", "ouoiiu", []*iface{&seat, nil, &surface, nil, nil, nil}},
		{"set_maximized", "?o", []*iface{&output}},
		{"set_title", "s", nil},
		{"set_class", "s", nil},
	},
	events: []message{
		{"ping", "u", nil},
		{"configure", "uii", nil},
		{"popup_done", "", nil}},
}
var compositor = iface{
	name: "wl_compositor", version: 3,
	methods: []message{
		{"create_surface", "n", []*iface{&surface}},
		{"create_region", "n", []*iface{&region}}},
	events: nil,
}

var region = iface{
	name: "wl_region", version: 1,
	methods: []message{
		{"destroy", "", nil},
		{"add", "iiii", nil},
		{"subtract", "iiii", nil}},
	events: nil,
}

var surface = iface{
	name: "wl_surface", version: 3,
	methods: []message{
		{"destroy", "", nil},
		{"attach", "?oii", []*iface{&buffer, nil, nil}},
		{"damage", "iiii", nil},
		{"frame", "n", []*iface{&callback}},
		{"set_opaque_region", "?o", []*iface{&region}},
		{"set_input_region", "?o", []*iface{&region}},
		{"commit", "", nil},
		{"set_buffer_transform", "2i", nil},
		{"set_buffer_scale", "3i", nil}},
	events: []message{
		{"enter", "o", []*iface{&output}},
		{"leave", "o", []*iface{&output}}},
}
var seat = iface{
	name: "wl_seat", version: 3,
	methods: []message{
		{"get_pointer", "n", []*iface{&pointer}},
		{"get_keyboard", "n", []*iface{&keyboard}},
		{"get_touch", "n", []*iface{&touch}}},
	events: []message{
		{"capabilities", "u", nil},
		{"name", "2s", nil}},
}
var pointer = iface{
	name: "wl_pointer", version: 3,
	methods: []message{
		{"set_cursor", "u?oii", []*iface{nil, &surface, nil, nil}},
		{"release", "3", nil}},
	events: []message{
		{"enter", "3uoff", []*iface{nil, &surface, nil, nil}},
		{"leave", "3uo", []*iface{nil, &surface}},
		{"motion", "3uff", nil},
		{"button", "3uuuu", nil},
		{"axis", "3uuf", nil}},
}
var keyboard = iface{
	name: "wl_keyboard", version: 3,
	methods: []message{
		{"release", "3", nil}},
	events: []message{
		{"keymap", "3uhu", nil},
		{"enter", "3uoa", []*iface{nil, &surface, nil}},
		{"leave", "3uo", []*iface{nil, &surface}},
		{"key", "3uuuu", nil},
		{"modifiers", "3uuuuu", nil}},
}
var touch = iface{
	name: "wl_touch", version: 3,
	methods: []message{
		{"release", "3", nil}},
	events: []message{
		{"down", "3uuoiff", []*iface{nil, nil, &surface, nil, nil, nil}},
		{"up", "3uui", nil},
		{"motion", "3uiff", nil},
		{"frame", "3", nil},
		{"cancel", "3", nil}},
}
var subcompositor = iface{
	name: "wl_subcompositor", version: 1,
	methods: []message{
		{"destroy", "", nil},
		{"get_subsurface", "noo", []*iface{&subsurface, &surface, &surface}}},
	events: nil,
}
var subsurface = iface{
	name: "wl_subsurface", version: 1,
	methods: []message{
		{"destroy", "", nil},
		{"set_position", "ii", nil},
		{"place_above", "o", []*iface{&surface}},
		{"place_below", "o", []*iface{&surface}},
		{"set_sync", "", nil},
		{"set_desync", "", nil}},
	events: nil,
}
