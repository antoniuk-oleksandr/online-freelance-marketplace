<script lang="ts">
    import BreadcrumbsLayout from "@/common-components/Breadcrumbs/BreadcrumbsLayout.svelte";
    import {Link} from "svelte-routing";
    import type {BreadcrumbsItem} from "@/types/BreadcrumbsItem.ts";

    type BreadcrumbsProps = {
        items: BreadcrumbsItem[]
    }

    const {items}: BreadcrumbsProps = $props();

</script>

<BreadcrumbsLayout>
    {#each items as item, index}
        <Link to={item.href}>
            <p
                    class="{item.active === false ? 'cursor-auto' : 'hover:underline'} duration-200 ease-out text-cyan-500 text-base"
            >{item.title}</p>
        </Link>
        {#if index < items.length - 1}
            <p class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary">/</p>
        {/if}
    {/each}
</BreadcrumbsLayout>