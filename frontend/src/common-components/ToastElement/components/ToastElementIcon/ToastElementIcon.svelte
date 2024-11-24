<script lang="ts">
    import Icon from "@iconify/svelte";
    import type {ToastElementStore} from "@/types/ToastElementStore.ts";

    const {type}: ToastElementStore = $props();
</script>

{#if type === 'error'}
    <Icon class="text-2xl text-red-500" icon="material-symbols:error-outline"/>
{:else if type === 'success'}
    <Icon class="text-2xl text-green-500" icon="material-symbols:check"/>
{:else if type === 'info'}
    <Icon class="text-2xl text-cyan-500" icon="material-symbols:info-outline"/>
{:else if type === 'warning'}
    <Icon class="text-2xl text-orange-500" icon="material-symbols:warning-outline"/>
{/if}