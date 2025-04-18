<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    let {children}: LayoutProps = $props();
</script>

<div
        class={"mx-auto px-6 w-full md:max-w-256 xl:max-w-wrapper flex"}
>
    {@render children()}
</div>
