<script lang="ts">
    import Icon from "@iconify/svelte";
    import {handleGoogleAuth, handleGoogleButtonClick} from "@/common-components/Sign/handlers.ts";
    import Spinner from "@/common-components/Spinner/Spinner.svelte";
    import type {SignProps} from "@/types/SignProps.ts";

    const {googleButtonText}: SignProps = $props();

    const clientId = import.meta.env.VITE_GOOGLE_CLIENT_ID;

    let loading = $state(false);
    const setLoading = (value: boolean) => loading = value;
</script>

<button
        disabled={loading}
        onclick={() => handleGoogleButtonClick(setLoading, clientId)}
        type="button"
        class="{loading ? 'opacity-70' : 'active:scale-95 hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover'} flex !h-12 bg-light-palette-background-block dark:bg-dark-palette-background-block font-semibold duration-200 ease-out items-center justify-center border border-light-palette-divider dark:border-dark-palette-divider rounded-md gap-x-2 w-full"
>
    {#if loading}
        <Spinner size="size-8" color="border-l-white"/>
    {:else}
        <Icon icon="flat-color-icons:google" width="24" height="24"/>
        <span>{googleButtonText}</span>
    {/if}
</button>
