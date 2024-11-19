<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    let {children}: LayoutProps = $props();

</script>

<div class="flex gap-x-2.5 items-center">
    {@render children()}
</div>