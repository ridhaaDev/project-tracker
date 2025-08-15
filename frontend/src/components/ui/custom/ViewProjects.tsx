import  { useNavigate } from "react-router";
import { Card, CardContent, CardHeader, CardTitle } from "../card";
import { NavigationMenuComponent } from "./NavigationMenu";
import { Page } from "./Page";
import { useQuery } from "@tanstack/react-query";

type Project = {
  id: string;
  name: string;
  description: string;
  start_date: string;
};

export const ViewProjects = () => {
  const navigate = useNavigate();
  const fetchProjects = async () => {
    const response = await fetch("http://localhost:3000/api/v1/projects");
    if (!response.ok) {
      throw new Error("Failed to fetch projects");
    }
    return await response.json();
  };

  const {
    data: projects,
    isLoading,
    error,
  } = useQuery({
    queryKey: ["projects"],
    queryFn: fetchProjects,
  });

  console.log("Projects:", projects);

  return (
    <>
      <NavigationMenuComponent />
      <Page>
        {isLoading && <p>Loading projects...</p>}
        {error && <p>Error loading projects: {error.message}</p>}
        {projects && (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 mt-6">
            {projects.map((project: Project) => (
              <Card 
                key={project.id} 
                onClick={() => navigate(`/view/projects/${project.id}`)}
                className="cursor-pointer hover:scale-105 hover:drop-shadow-2xl transition-transform">
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
            ))}
          </div>
        )}

        {/* Add your project viewing logic here */}
      </Page>
    </>
  );
};
