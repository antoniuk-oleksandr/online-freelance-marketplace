<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type InputPasswordIconLayoutProps = LayoutProps & {
        showPassword: boolean,
        setShowPassword: (value: boolean) => void,
    }

    const {showPassword, setShowPassword, children}: InputPasswordIconLayoutProps = $props();
</script>

<button
        type="button"
        onclick={() => setShowPassword(!showPassword)}
        class="absolute text-light-palette-divider dark:text-dark-palette-divider cursor-pointer right-4 top-[14px] text-xl ring-light-palette-action-hover dark:ring-dark-palette-action-hover hover:ring-8 duration-200 ease-out hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover rounded-full">
    {@render children()}
</button>
