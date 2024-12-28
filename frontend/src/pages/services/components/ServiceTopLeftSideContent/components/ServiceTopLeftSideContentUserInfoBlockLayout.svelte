<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props()
</script>

<div class="grid place-items-center sm:block">
    {@render children()}
</div>