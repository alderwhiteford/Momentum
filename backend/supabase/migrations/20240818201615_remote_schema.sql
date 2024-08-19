create type "public"."oauthProvider" as enum ('google');

alter table "public"."users" alter column "provider" set data type "oauthProvider" using "provider"::text::"oauthProvider";

alter table "public"."users" enable row level security;

drop type "public"."oauthprovider";

create policy "Enable delete for users based on user_id"
on "public"."users"
as permissive
for delete
to public
using ((( SELECT auth.uid() AS uid) = user_id));


create policy "Enable read for users based on id"
on "public"."users"
as permissive
for select
to authenticated
using (( SELECT (auth.uid() = users.user_id)));


create policy "Enable update for users based on id"
on "public"."users"
as permissive
for update
to public
using ((( SELECT auth.uid() AS uid) = user_id))
with check ((( SELECT auth.uid() AS uid) = user_id));



