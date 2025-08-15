import React from "react";
import { Input } from "../input";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "../card";
import { z } from "zod";
import { toast } from "sonner";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../form";
import { Button } from "../button";
import { ButtonGroup } from "../button-group";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Textarea } from "../textarea";
import { Page } from "./Page";
import { NavigationMenuComponent } from "./NavigationMenu";

const FormSchema = z.object({
  username: z.string().min(2, {
    message: "Username must be at least 2 characters.",
  }),
});

export const AddProject = () => {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      username: "",
    },
  });

  function onSubmit(data: z.infer<typeof FormSchema>) {
    console.log(data);
    toast("You submitted the following values", {
      description: (
        <pre className="mt-2 w-[320px] rounded-md bg-neutral-950 p-4">
          <code className="text-white">{JSON.stringify(data, null, 2)}</code>
        </pre>
      ),
    });
  }

  return (
    <>
      <NavigationMenuComponent />
      <Page>
        <div className="flex flex-col gap-6 lg:min-w-[400px]">
          <Card>
            <CardHeader>
              <CardTitle>Add project details</CardTitle>
              <CardDescription>
                You can change this any time in the project setting
              </CardDescription>
            </CardHeader>
            <CardContent>
              <Form {...form}>
                <form
                  onSubmit={form.handleSubmit(onSubmit)}
                  className="w-2/3 space-y-6"
                >
                  <FormField
                    control={form.control}
                    name="username"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Project Name</FormLabel>
                        <FormControl>
                          <Input
                            placeholder="Try a team name, project or milestone"
                            {...field}
                            className="lg:w-66 md:w-max"
                          />
                        </FormControl>

                        <FormLabel className="mt-6">
                          Project Description
                        </FormLabel>
                        <FormControl>
                          <Textarea placeholder="Enter your project description" />
                        </FormControl>

                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <Button type="submit">Submit</Button>
                </form>
              </Form>
            </CardContent>
            <CardFooter>
              <ButtonGroup></ButtonGroup>
            </CardFooter>
          </Card>
        </div>
      </Page>
    </>
  );
};
