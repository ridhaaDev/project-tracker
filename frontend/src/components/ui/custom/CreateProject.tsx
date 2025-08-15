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
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "react-router";
// import { useNavigate } from "react-router";

const FormSchema = z.object({
  name: z.string().min(2, {
    message: "Project name must be at least 2 characters.",
  }),
  description: z.string(),
});

export const AddProject = () => {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      name: "",
      description: "",
    },
  });

  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const addProjectMutation = useMutation({
    mutationFn: async (data: z.infer<typeof FormSchema>) => {
      const response = await fetch(
        "http://localhost:3000/api/v1/projects/create",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(data),
        }
      );
      if (!response.ok) {
        throw new Error("Failed to add project");
      }
      return response.json();
    },
    onSuccess: () => {
      toast.success("Project added successfully!");
      queryClient.invalidateQueries({ queryKey: ["projects"] });
      navigate("/view/projects");
    },
    onError: (error: { message?: string }) => {
      toast.error(error.message || "Failed to add project");
    },
  });

  function onSubmit(data: z.infer<typeof FormSchema>) {
    addProjectMutation.mutate(data);

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
                    name="name"
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
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <FormField
                    control={form.control}
                    name="description"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Project Description</FormLabel>
                        <FormControl>
                          <Textarea
                            placeholder="Something to describe your project"
                            {...field}
                            className="lg:w-66 md:w-max"
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <Button type="submit">Start Project</Button>
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
