<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type StepperLayoutProps = LayoutProps & {
        styles?: string,
    }

    const {children, styles} : StepperLayoutProps = $props();
</script>

<div class="{styles} flex-wrap flex h-fit gap-x-3 text-lg items-center font-semibold">
    {@render children()}
</div>