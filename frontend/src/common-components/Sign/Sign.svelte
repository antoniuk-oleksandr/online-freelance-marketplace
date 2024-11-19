<script lang="ts">
    import type {SignProps} from "@/types/SignProps.ts";
    import SignLayout from "@/common-components/Sign/SignLayout.svelte";
    import LeftSignSide from "@/common-components/Sign/components/LeftSignSide/LeftSignSide.svelte";
    import RightSignSide from "@/common-components/Sign/components/RightSignSide/RightSignSide.svelte";

    const signProps: SignProps = $props();
</script>

<SignLayout>
    <LeftSignSide {...signProps}/>
    <RightSignSide/>
</SignLayout>
