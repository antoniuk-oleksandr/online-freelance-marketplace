<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="flex flex-row items-center lg:flex-col xl:flex-row justify-between">
    {@render children()}
</div>
