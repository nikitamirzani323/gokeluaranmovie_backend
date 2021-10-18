<script>
    import { link } from "svelte-spa-router";
    async function handleLogout() {
        localStorage.clear();
        window.location.href = "/";
    }
</script>
<!-- Navbar -->
<nav class="navbar navbar-expand-lg bg-dark navbar-dark fixed-top">
  <!-- Container wrapper -->
    <div class="container-fluid">

        <!-- Toggle button -->
        <button class="navbar-toggler" type="button" data-mdb-toggle="collapse" data-mdb-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <i class="fas fa-bars"></i>
        </button>

        <!-- Collapsible wrapper -->
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                <!-- Link -->
                <li class="nav-item">
                    <a class="nav-link" use:link href="/">DASHBOARD</a>
                </li>
                <li class="nav-item dropdown">
                    <a class="nav-link dropdown-toggle"  id="navbarDropdown" role="button" data-mdb-toggle="dropdown" aria-expanded="false">
                        MASTER
                    </a>
                    <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
                        <li>
                            <a class="dropdown-item" use:link href="/admin">ADMIN</a>
                        </li>
                        <li>
                            <a class="dropdown-item" use:link href="/adminrule">ADMIN RULE</a>
                        </li>
                    </ul>
                </li>

            </ul>

            <!-- Icons -->
            <ul class="navbar-nav d-flex flex-row me-1">
                <li class="nav-item me-3 me-lg-0">
                    <button
                        on:click={() => {
                            handleLogout();
                        }} 
                        type="button" class="btn btn-warning btn-sm">LOGOUT</button>
                </li>
            </ul>
        </div>
    </div>
</nav>
<!-- Navbar -->