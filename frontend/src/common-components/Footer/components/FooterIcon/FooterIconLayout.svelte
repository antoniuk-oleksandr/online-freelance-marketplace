<script>
    export let link;
</script>


<button
        type="button"
        on:click={() => window.open(link, "_blank")}
        class={"cursor-pointer ring-light-palette-action-hover dark:ring-dark-palette-action-hover hover:ring-8 duration-200 ease-out hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover rounded-full"}
>
    <slot/>
</button>