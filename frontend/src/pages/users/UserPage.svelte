<script lang="ts">
    import type {User} from "@/types/User";
    import UserPageLayout from "@/pages/users/UserPageLayout.svelte";
    import UserAboutBlock from "@/pages/users/components/UserAboutBlock/UserAboutBlock.svelte";
    import UserSkillsBlock from "@/pages/users/components/UserSkillsBlock/UserSkillsBlock.svelte";
    import UserInfoBlock from "@/pages/users/components/UserInfoBlock/UserInfoBlock.svelte";
    import UserByIdServicesBlock from "@/pages/users/components/UserByIdServicesBlock/UserByIdServicesBlock.svelte";
    import UserByIdReviewBlock from "@/pages/users/components/UserByIdReviewBlock/UserByIdReviewBlock.svelte";
    import NotFound from "@/common-components/NotFound/NotFound.svelte";
    import {tryToGetUserById} from "@/pages/users/helpers.ts";

    type UserPageProps = {
        id: string,
    }

    let {id}: UserPageProps = $props();

    let user = $state<User | null | undefined>();
    const setUser = (newUser: User | null | undefined) => user = newUser;

    $effect(() => {
        tryToGetUserById(id, setUser);
    })
</script>

{#if user === null}
    <NotFound/>
{:else if user !== undefined}
    <UserPageLayout>
        <UserInfoBlock size={"large"} user={user}/>
        <UserAboutBlock about={user.about}/>
        <UserSkillsBlock skills={user.skills}/>
        <UserByIdServicesBlock services={user.services}/>
        <UserByIdReviewBlock showServices={true} reviews={user.reviews}/>
    </UserPageLayout>
{/if}