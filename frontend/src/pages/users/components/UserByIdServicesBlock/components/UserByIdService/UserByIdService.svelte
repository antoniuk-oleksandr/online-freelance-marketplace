<script lang="ts">
    import type {UserService} from "@/types/UserService.ts";
    import UserByIdServiceLayout
        from "@/pages/users/components/UserByIdServicesBlock/components/UserByIdService/UserByIdServiceLayout.svelte";
    import UserByIdServiceImage
        from "@/pages/users/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServiceImage.svelte";
    import UserByIdServiceTitle
        from "@/pages/users/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServiceTitle.svelte";
    import UserByIdServiceInfo
        from "@/pages/users/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServiceInfo/UserByIdServiceInfo.svelte";
    import UserByIdServicePrice
        from "@/pages/users/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServicePrice.svelte";

    type UserByIdServiceProps = {
        service: UserService,
    }

    const {service}: UserByIdServiceProps = $props();
</script>

<UserByIdServiceLayout {...service}>
    <UserByIdServiceImage {...service}/>
    <UserByIdServiceTitle {...service}/>
    <UserByIdServiceInfo {...service}/>
    <UserByIdServicePrice {...service}/>
</UserByIdServiceLayout>