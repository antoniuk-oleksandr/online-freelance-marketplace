<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import type {ToastElementStore} from "@/types/ToastElementStore.ts";

    const {children, show, exitAnimation}: LayoutProps & ToastElementStore = $props();
</script>

<div class="fixed left-6 bottom-6">
    <div class="{exitAnimation ? 'animate-toastExit' : 'animate-toast'} rounded-md ring-1 ring-light-palette-divider dark:ring-dark-palette-divider p-4 bg-light-palette-background-block dark:bg-dark-palette-background-block flex items-center gap-x-2">
        {@render children()}
    </div>
</div>