<script lang="ts">
    import type {BreadcrumbsItem} from "@/types/BreadcrumbsItem.ts";
    import Breadcrumbs from "@/common-components/Breadcrumbs/Breadcrumbs.svelte";
    import type {Service} from "@/types/Service.ts";

    type ServiceBreadcrumbsProps = {
        service: Service
    }

    const {service}: ServiceBreadcrumbsProps = $props();

    const items: BreadcrumbsItem[] = [
        {title: "Home", href: "/"},
        {title: service.category.name, href: `/search?category=${service.category.id}`},
        {title: service.title, href: `/services/${service.id}`, active: false}
    ];
</script>

<Breadcrumbs items={items}/>