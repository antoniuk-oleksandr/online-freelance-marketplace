<script lang="ts">
    type CardInputErrorProps = {
        error?: string,
    }

    let {error}: CardInputErrorProps = $props();
</script>

{#if error}
    <p class="text-red-500">{error}</p>
{/if}