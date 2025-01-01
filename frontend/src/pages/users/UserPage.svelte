<script lang="ts">
    import UserPageLayout from "@/pages/users/UserPageLayout.svelte";
    import UserAboutBlock from "@/pages/users/components/UserAboutBlock/UserAboutBlock.svelte";
    import UserSkillsBlock from "@/pages/users/components/UserSkillsBlock/UserSkillsBlock.svelte";
    import UserInfoBlock from "@/pages/users/components/UserInfoBlock/UserInfoBlock.svelte";
    import UserByIdServicesBlock from "@/pages/users/components/UserByIdServicesBlock/UserByIdServicesBlock.svelte";
    import UserByIdReviewBlock from "@/pages/users/components/UserByIdReviewBlock/UserByIdReviewBlock.svelte";
    import NotFound from "@/common-components/NotFound/NotFound.svelte";
    import {getMoreReviewsByUserId, getMoreServicesByUserId, tryToGetUserById} from "@/pages/users/helpers.ts";
    import {userByIdStore} from "@/pages/users/stores/user-by-id-store.ts";
    import type {UserByIdData} from "@/types/GetUserByIdData.ts";

    type UserPageProps = {
        id: string,
    }

    let {id}: UserPageProps = $props();

    let userData = $state<UserByIdData | null | undefined>();
    userByIdStore.subscribe((val) => userData = val);

    tryToGetUserById(id);
</script>

{#if userData === null}
    <NotFound/>
{:else if userData !== undefined}
    <UserPageLayout>
        <UserInfoBlock size={"large"} user={userData.user}/>
        <UserAboutBlock about={userData.user.about}/>
        <UserSkillsBlock skills={userData.user.skills}/>
        <UserByIdServicesBlock
                showMoreServicesButtonAction={() => getMoreServicesByUserId(userData)}
                hasMore={userData.hasMoreServices}
                services={userData.user.services}/>
        <UserByIdReviewBlock
                showMoreReviewsButtonAction={() => getMoreReviewsByUserId(userData)}
                hasMore={userData.hasMoreReviews}
                reviews={userData.user.reviews}
                showServices={true}
        />
    </UserPageLayout>
{/if}