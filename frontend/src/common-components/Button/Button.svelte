<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import Spinner from "@/common-components/Spinner/Spinner.svelte";

    type MyButtonProps = LayoutProps & {
        clickAction?: () => void,
        styles?: string,
        type?: "button" | "submit" | "reset"
        loading?: boolean
    };

    const {children, clickAction, styles, type, loading}: MyButtonProps = $props();
</script>

<button
        disabled={loading}
        type={type}
        class="{styles} {loading ? 'opacity-70' : 'hover:!bg-cyan-400 active:!scale-95'} !bg-cyan-500 text-white px-4 rounded-md !h-12 !grid !place-items-center !duration-200 font-semibold !ease-out"
        onclick={clickAction}
>
    {#if loading}
        <Spinner size="size-8" color="border-l-white"/>
    {:else}
        {@render children()}
    {/if}
</button>