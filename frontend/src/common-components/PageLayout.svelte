<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    let {children}: LayoutProps = $props();
</script>

<div
        class={"mx-auto px-4 w-full md:max-w-256 xl:max-w-320 flex"}
>
    {@render children()}
</div>