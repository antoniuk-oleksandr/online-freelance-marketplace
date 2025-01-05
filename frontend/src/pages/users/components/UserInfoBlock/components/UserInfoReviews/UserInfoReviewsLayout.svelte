<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type UserInfoReviewsLayoutProps = LayoutProps & {
        size?: "small" | "large",
    };

    const {children, size}: UserInfoReviewsLayoutProps = $props();
</script>


<div
        class={`flex items-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary gap-x-1
        ${size === "small" ? "text-base" : "text-lg"}`}
>
    {@render children()}
</div>