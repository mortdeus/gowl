package client

type Proxy interface {
	Marshal(opcode ...uint)
	Display() Display
	Events() EventQueue
}

func CreateProxy(factory Proxy, typ meta) Proxy {
	switch typ.name {
	case "wl_display":
	case "wl_registry":
	case "wl_callback":
	case "wl_shm_pool":
	case "wl_shm":
	case "wl_buffer":
	case "wl_output":
	case "wl_data_offer":
	case "wl_data_source":
	case "wl_data_device":
	case "wl_data_device_manager":
	case "wl_shell":
	case "wl_shell_surface":
	case "wl_compositor":
	case "wl_region":
	case "wl_surface":
	case "wl_seat":
	case "wl_pointer":
	case "wl_keyboard":
	case "wl_touch":
	case "wl_subcompositor":
	case "wl_subsurface":

	default:
		//TODO(mortdeus): check against 3rd party dynamically
		//registered extensions.
	}
}
