<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";


    type MyButtonProps = LayoutProps & {
        clickAction?: () => void,
        styles?: string,
        type?: "button" | "submit" | "reset"
    };

    const {children, clickAction, styles, type}: MyButtonProps = $props();
</script>

<button
        type={type}
        class="{styles} text-white rounded-md !h-12 !grid !place-items-center !bg-cyan-500 hover:!bg-cyan-400 active:!scale-95 !duration-200 font-semibold !ease-out"
        onclick={clickAction}
>
    {@render children()}
</button>