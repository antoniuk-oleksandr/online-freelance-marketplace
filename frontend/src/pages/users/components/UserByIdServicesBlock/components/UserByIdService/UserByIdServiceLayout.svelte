<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {Link} from "svelte-routing";
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import {getHoverClass} from "@/utils/utils.ts";

    type UserByIdServiceLayoutProps = LayoutProps & {
        id: number,
        size?: "small" | "normal"
    }

    const {children, id, size}: UserByIdServiceLayoutProps = $props();
    const hoverClass = getHoverClass();
</script>

<Link to={`/services/${id}`}>
    <PaperElement styles="service-element {hoverClass} {size === 'small' ? '!p-3' : ''}">
        {@render children()}
    </PaperElement>
</Link>