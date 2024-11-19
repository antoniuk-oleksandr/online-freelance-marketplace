<script lang="ts">
    import type {UserService} from "@/types/UserService.ts";
    import UserByIdServiceInfoLayout
        from "@/pages/users/components/UserByIdServicesBlock/components/UserByIdService/components/UserByIdServiceInfo/UserByIdServiceInfoLayout.svelte";
    import Icon from "@iconify/svelte";

    const {rating, reviewsCount}: UserService = $props();
</script>

<UserByIdServiceInfoLayout>
    <Icon class="text-yellow-500 mr-1" icon="mingcute:star-fill"/>
    <span class="font-semibold vertical-align-middle">{rating}</span>
    <span class="ml-1  vertical-align-middle">({reviewsCount})</span>
</UserByIdServiceInfoLayout>