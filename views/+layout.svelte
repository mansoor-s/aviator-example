<script lang="js">

    //import custom material UI theme from CSS file
    import './svelteTheme.css';

    export let Title;
    console.log("title is: " + Title);
</script>


<svelte:head>
    <title>{Title}</title>


    <!-- <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css"> -->

    <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/icon?family=Material+Icons"
    />
    <!-- Roboto -->
    <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,600,700"
    />
    <!-- Roboto Mono -->
    <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/css?family=Roboto+Mono"
    />

    <link
        rel="stylesheet"
        href="https://cdn.jsdelivr.net/npm/svelte-material-ui@6.0.0/bare.min.css"
    />
</svelte:head>

<h1>{Title}</h1>
<br/>
<br/>

<slot />
