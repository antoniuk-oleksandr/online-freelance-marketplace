<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    let {children}: LayoutProps = $props();

</script>

<div class="flex gap-x-1 items-center flex-wrap">
    {@render children()}
</div>