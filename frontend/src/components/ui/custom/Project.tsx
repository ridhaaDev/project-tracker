import { Label } from "../label";
import { Input } from "../input";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "../card";
import { Button } from "../button";
import { ButtonGroup } from "../button-group";

export const AddProject = () => {
  return (
      <div className="flex flex-col gap-6 min-w-[400px]">
        <Card>
          <CardHeader>
            <CardTitle>Add project details</CardTitle>
            <CardDescription>
              You can change this any time in the project setting
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div>
              <Label htmlFor="project-name">Name</Label>
              <Input
                type="text"
                id="project-name"
                placeholder="Try a team name, project or milestone"
                required
              />
            </div>
          </CardContent>
          <CardFooter>
            <ButtonGroup>
              <Button className="cursor-pointer">Create project</Button>
              <Button className="cursor-pointer" variant={"outline"}>Cancel</Button>
            </ButtonGroup>
          </CardFooter>
        </Card>
      </div>
  );
};
