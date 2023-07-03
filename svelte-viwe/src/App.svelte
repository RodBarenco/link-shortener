<script>
    import List from "./components/List.svelte";
    import New from "./components/New.svelte";
</script>

<main>
	<container class="principal">
    <h1>&nbsp;Goly -- The Go URL Shortener! &nbsp; <New/></h1>
    <List />
    </container>
</main>

<style>
    .principal {
		display: flex;
        flex-direction: column;
        justify-content: center; 
        align-items: center; 
        height: 100%; 
    }
</style>