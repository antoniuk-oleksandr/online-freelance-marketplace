<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type SelectMenuLayoutProps = LayoutProps & {
        isOpen: boolean,
    };

    const {children, isOpen}: SelectMenuLayoutProps = $props();
</script>

<div class="absolute">
    <div
            class="{isOpen ? 'animate-drop-down' : 'animate-drop-up'} Select-menu static flex flex-col shadow-xl gap-y-2 p-1 mt-2 rounded-lg right-0 whitespace-nowrap bg-light-palette-background-block dark:bg-dark-palette-background-block"
    >
        {@render children()}
    </div>
</div>
