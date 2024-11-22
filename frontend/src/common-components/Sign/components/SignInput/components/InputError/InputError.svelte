<script lang="ts">
    type InputErrorProps = {
        wasSubmitted: boolean,
        error: string | null,
    }

    const {wasSubmitted, error}: InputErrorProps = $props();
</script>

{#if wasSubmitted && error}
    <span class="text-red-500">{error}</span>
{/if}