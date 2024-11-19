<script lang="ts">
    import ServiceBottomLeftSideContentLayout
        from "@/pages/services/components/ServiceBottomLeftSideContent/ServiceBottomLeftSideContentLayout.svelte";
    import ServiceAboutBlock from "@/pages/services/components/ServiceAboutBlock/ServiceAboutBlock.svelte";
    import UserByIdReviewBlock from "@/pages/users/components/UserByIdReviewBlock/UserByIdReviewBlock.svelte";
    import type {Service} from "@/types/Service.ts";

    const serviceProps: Service = $props();
</script>

<ServiceBottomLeftSideContentLayout>
    <ServiceAboutBlock {...serviceProps}/>
    <UserByIdReviewBlock showServices={false} {...serviceProps}/>
</ServiceBottomLeftSideContentLayout>