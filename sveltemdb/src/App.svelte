<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navbar from "./components/Navbar.svelte";
	import Company from "./pages/company/Company.svelte";
	import Admin from "./pages/admin/Admin.svelte";
	import Login from "./pages/Login.svelte";
	import NotFound from "./pages/Notfound.svelte";
	export let table_header_font
	export let table_body_font
	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
	if (token == "" || token == null) {
		routes = {
			"/": wrap({
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
		routes = {
			"/": wrap({
				component: Company,
			}),
			"/admin": wrap({
				component: Admin,
				props:{
					table_header_font:table_header_font,
					table_body_font:table_body_font
				}
			}),
			"*": NotFound,
		};
	}
</script>

{#if isNav}
	<Navbar />
{/if}
<Router {routes} />
