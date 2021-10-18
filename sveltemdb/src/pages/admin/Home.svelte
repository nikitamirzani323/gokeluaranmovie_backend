<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";

	export let table_header_font
	export let table_body_font
	export let listAdmin = []
	export let listAdminrule = []
	export let totalrecord = 0
	let title_page = "Admin"
    let sData = "";
    let myModal = ""
    let css_loader = "display: none;";
    let msgloader = "";
    const schema = yup.object().shape({
        username: yup.string().required("The Username is required").
                    matches(/^[a-zA-z0-9]+$/, "Username must Character A-Z or a-z or 1-9 ").
                    min(3,"The Username minimal 3 Character").
                    max(30,"The Username mmaximal 30 Character"),
        password: yup.string().required("The Password is required").
                    min(4,"The Password minimal 3 Character").
                    max(30,"The Password mmaximal 30 Character"),
        rule: yup.string().required("The Rule is required")
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            username: "",
            password: "",
            rule:""
        },
        validationSchema: schema,
        onSubmit:(values) => {
            handleSave(values.username,values.password,value.rule)
        }
    })
    const NewData = () => {
        sData = "New"
        myModal = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal.show();
    };
    const RefreshData = () => {
        sData = "New"
        myModal.hide();
    };
    async function handleSave(username,password,rule) {
        const res = await fetch("/api/saveadmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                idruleadmin: parseInt(rule),
                page:"ADMIN-SAVE",
                username: username,
                password: password,
                nama: name,
                status: "ACTIVE",
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            msgloader = json.message;
            $form.username = ""
            $form.password = ""
            $form.rule = ""
        } else if(json.status == 403){
            alert(json.message)
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
        
    }
    $:{
        if ($errors.username || $errors.password || $errors.rule){
            alert($errors.username+"\n"+$errors.password+"\n"+$errors.rule)
            $form.username = ""
            $form.password = ""
            $form.rule = ""
        }
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
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
                    {#if totalrecord > 0}
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
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;"><i class="bi bi-pencil"></i></td>
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
                    {:else}
                        <center>
                            <Loader />
                        </center>
                    {/if}
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentry"
	modal_size="modal-dialog-centered"
	modal_title="{title_page+"/"+sData}"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Rule</label>
			<select
                on:change="{handleChange}"
                bind:value={$form.rule} 
                class="form-control required">
                <option value="">--Select--</option>
                {#each listAdminrule as rec }
                    <option value="{rec.adminrule_idruleadmin}">{rec.adminrule_idruleadmin}</option>
                {/each}
            </select>
		</div>
		<div class="mb-3">
            <label for="exampleForm" class="form-label">Username</label>
            <input
                on:change="{handleChange}"
                bind:value={$form.username}
                invalid={$errors.username.length > 0}
                type="text"
                class="form-control required"
                placeholder="Username"
                aria-label="Username"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Password</label>
            <input
                on:change="{handleChange}"
                bind:value={$form.password}
                invalid={$errors.password.length > 0}
                type="password"
                class="form-control required"
                placeholder="Password"
                aria-label="Password"/>
        </div>
	</slot:template>
	<slot:template slot="footer">
		<button
            on:click={() => {
                handleSubmit();
            }} 
            type="button" class="btn btn-dark">Save</button>
	</slot:template>
</Modal>