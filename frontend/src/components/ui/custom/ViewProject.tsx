import { Link, useLocation } from "react-router";
import { NavigationMenuComponent } from "./NavigationMenu";
import { FullPage } from "./Page";
import { useQuery } from "@tanstack/react-query";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "../breadcrumb";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "../table";
import { Button } from "../button";
import { Input } from "../input";

// type Project = {
//   id: string;
//   name: string;
//   description: string;
//   start_date: string;
// };

type Ticket = {
  id: string;
  name: string;
  description: string;
  status: string;
  assignee: string;
};

export const ViewProject = () => {
  const location = useLocation();

  const fetchProject = async () => {
    const projectId = location.pathname.split("/").pop();

    const response = await fetch(
      `http://localhost:3000/api/v1/projects/${projectId}`
    );
    if (!response.ok) {
      throw new Error("Failed to fetch project");
    }
    return response.json();
  };

  const {
    data: project,
    isLoading,
    error,
  } = useQuery({
    queryKey: ["project"],
    queryFn: fetchProject,
  });

  const fetchTickets = async () => {
    const projectId = location.pathname.split("/").pop();
    const response = await fetch(
      `http://localhost:3000/api/v1/projects/${projectId}/tickets`
    );
    if (!response.ok) {
      throw new Error("Failed to fetch tickets");
    }
    return response.json();
  };

  const {
    data: tickets,
    isLoading: isTicketsLoading,
    error: ticketsError,
  } = useQuery({
    queryKey: ["tickets", location.pathname],
    queryFn: fetchTickets,
  });

  return (
    <>
      <NavigationMenuComponent />

      <FullPage>
        {isLoading && <p>Loading projects...</p>}
        {error && <p>Error loading projects: {error.message}</p>}

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4 mt-6">
          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem>
                <BreadcrumbLink href="/view/projects">Projects</BreadcrumbLink>
              </BreadcrumbItem>
              <BreadcrumbSeparator />
              <BreadcrumbItem>
                <BreadcrumbPage>{project?.name}</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
            <div className="flex flex-col gap-4 mt-6">
              <div className="col-span-1">
                <h2 className="text-sm font-semibold">Project Details</h2>
                <p className="text-gray-600">{project?.description}</p>
              </div>
              <div className="col-span-1">
                <h2 className="text-sm font-semibold">Project Info</h2>
                <p className="text-gray-600">
                  Created: {new Date(project?.start_date).toLocaleDateString()}
                </p>
              </div>
            </div>
          </Breadcrumb>

          <div className="col-span-3 w-full">
            <h2 className="text-lg font-semibold mb-4">Sprints</h2>

            <Button>Create Sprint</Button>

            <Table className="">
              {/* <TableCaption>A list of your recent invoices.</TableCaption> */}
              <TableHeader className="">
                <TableRow>
                  <TableHead>Sprint name</TableHead>
                  <TableHead>Description</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead>Assignee</TableHead>
                  <TableHead>Points</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {isTicketsLoading && (
                  <TableRow>
                    <TableCell colSpan={5}>Loading tickets...</TableCell>
                  </TableRow>
                )}
                {ticketsError && (
                  <TableRow>
                    <TableCell colSpan={5}>
                      Error loading tickets: {ticketsError.message}
                    </TableCell>
                  </TableRow>
                )}
                {tickets?.map((ticket: Ticket) => (
                  <TableRow key={ticket.id}>
                    <TableCell className="font-medium">{ticket.name}</TableCell>
                    <TableCell>{ticket.description}</TableCell>
                    <TableCell>{ticket.status}</TableCell>
                    <TableCell>{ticket.assignee}</TableCell>

                    <TableRow>
                      <TableCell className="text-right">
                        <Link
                          className="hover:underline"
                          to={{ pathname: "/" }}
                        >
                          View Details
                        </Link>
                      </TableCell>
                    </TableRow>
                  </TableRow>
                ))}

                <TableRow>
                  <TableCell>
                    <Input value={""} />
                  </TableCell>
                  <TableCell>
                    <Input value={""} />
                  </TableCell>
                  <TableCell>
                    <Input value={""} />
                  </TableCell>
                  <TableCell>
                    <Input value={""} />
                  </TableCell>
    
                </TableRow>
              </TableBody>
            </Table>
          </div>

          <div className="col-span-3 col-start-2 w-full">
            <h2 className="text-lg font-semibold mb-4">Backlog</h2>
          </div>
        </div>

        {/* {project && (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 mt-6">
            <Card
              key={project.id}
              className="cursor-pointer hover:scale-105 hover:drop-shadow-2xl transition-transform"
            >
              <CardHeader>
                <CardTitle>{project.name}</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-gray-600 mb-2">{project.description}</p>
                <span className="text-sm text-gray-500">
                  Created: {new Date(project.start_date).toLocaleDateString()}
                </span>
              </CardContent>
            </Card>
          </div>
        )} */}
        {/* Add your project viewing logic here */}
      </FullPage>
    </>
  );
};
