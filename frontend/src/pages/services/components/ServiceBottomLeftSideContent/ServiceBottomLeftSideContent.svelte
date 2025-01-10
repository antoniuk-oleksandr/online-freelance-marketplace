<script lang="ts">
    import ServiceBottomLeftSideContentLayout
        from "@/pages/services/components/ServiceBottomLeftSideContent/ServiceBottomLeftSideContentLayout.svelte";
    import ServiceAboutBlock from "@/pages/services/components/ServiceAboutBlock/ServiceAboutBlock.svelte";
    import UserByIdReviewBlock from "@/pages/users/components/UserByIdReviewBlock/UserByIdReviewBlock.svelte";
    import type {UpdateFunc} from "@/types/UpdateFunc.ts";
    import type {GetServiceByUserIdRequestResponse} from "@/types/GetServiceByIdRequestResponse.ts";
    import {handleShowMoreFreelancesButtonClick} from "@/pages/services/handlers.ts";

    type ServiceBottomLeftSideContentProps = {
        setServiceResponse: UpdateFunc<GetServiceByUserIdRequestResponse | undefined>,
        serviceResponse: GetServiceByUserIdRequestResponse
    };

    const serviceProps: ServiceBottomLeftSideContentProps = $props();
</script>

{#if serviceProps.serviceResponse.status === 200}
    <ServiceBottomLeftSideContentLayout>
        <ServiceAboutBlock {...serviceProps.serviceResponse.data.service}/>
        <UserByIdReviewBlock
                hasMore={serviceProps.serviceResponse.data.hasMoreReviews}
                showMoreReviewsButtonAction={() => handleShowMoreFreelancesButtonClick(serviceProps.setServiceResponse)}
                showServices={false}
                {...serviceProps.serviceResponse.data.service}/>
    </ServiceBottomLeftSideContentLayout>
{/if}