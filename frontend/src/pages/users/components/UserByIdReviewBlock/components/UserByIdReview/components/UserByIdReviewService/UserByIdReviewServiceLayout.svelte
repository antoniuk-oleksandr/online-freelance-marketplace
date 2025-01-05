<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {Link} from "svelte-routing";
    import {getHoverClass} from "@/utils/utils.ts";

    type UserByIdReviewServiceLayoutPros = LayoutProps & {
        id: number
    };

    const {children, id}: UserByIdReviewServiceLayoutPros = $props();
    const hoverClass = getHoverClass();
</script>

<Link to={`/services/${id}`}>
    <div class="{hoverClass} py-2 px-4 ring-1 ring-light-palette-divider dark:ring-dark-palette-divider shadow rounded-xl flex gap-3 items-center w-full md:w-fit">
        {@render children()}
    </div>
</Link>
