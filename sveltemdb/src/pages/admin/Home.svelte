<script>
    import Panel from "../../components/Panel.svelte";
	import Modal from "../../components/Modal.svelte";

	export let table_header_font
	export let table_body_font
	export let listAdmin = []
	export let listAdminrule = []
	export let totalrecord = 0
	let title_page = "Home"
    let sData = "";
    let myModal = ""
    const NewData = () => {
        sData = "New"
        myModal = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal.show();
    };
    const RefreshData = () => {
        sData = "New"
        myModal.hide();
    };
</script>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <button
                on:click={() => {
                    NewData();
                }}
                type="button"
                class="btn btn-dark btn-sm">
                NEW
            </button>
            <button
                on:click={() => {
                    RefreshData();
                }}
                type="button"
                class="btn btn-primary btn-sm">
                Refresh
            </button>
            <Panel
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-body">
                    <table class="table table-striped table-hover">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">RULE</th>
                                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TIMEZONE</th>
                                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">JOIN DATE</th>
                                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">LAST LOGIN</th>
                                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">LAST IPADDRESS</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each listAdmin as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;"><i class="far fa-edit"></i></td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};{rec.admin_statuscss}">{rec.admin_status}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_no}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_username}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_nama}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_rule}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_timezone}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_joindate}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_lastlogin}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_lastipaddres}</td>
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentry"
	modal_size="modal-dialog-centered"
	modal_title="{sData}"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
			<select class="form-control">
                {#each listAdminrule as rec }
                    <option value="{rec.adminrule_idruleadmin}">{rec.adminrule_idruleadmin}</option>
                {/each}
            </select>
		</div>
		<div class="mb-3">
			<input type="type" class="form-control" id="username" placeholder="Username">
		</div>
		<div class="mb-3">
			<input type="password" class="form-control" id="password" placeholder="Password">
		</div>
	</slot:template>
	<slot:template slot="footer">
		<button
            on:click={() => {
                RefreshData();
            }} 
            type="button" class="btn btn-dark">Save</button>
	</slot:template>
</Modal>