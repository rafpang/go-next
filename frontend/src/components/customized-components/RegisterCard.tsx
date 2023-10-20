import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function RegisterCard() {
  return (
    <Card className="w-[400px] md:w-[50vw] shadow-lg mx-auto">
      <CardHeader className="flex justify-center items-center">
        <CardTitle>Register</CardTitle>
        <CardDescription>Sign into the App</CardDescription>
      </CardHeader>
      <CardContent>
        <form>
          <div className="grid w-full items-center gap-4">
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="name">Username</Label>
              <Input id="name" placeholder="Enter your username" />
            </div>
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="name">Password</Label>
              <Input id="name" placeholder="Enter your password" />
            </div>
          </div>
        </form>
      </CardContent>
      <CardFooter className="flex flex-col gap-5">
        <Button variant="outline" className="w-[70%]">
          Log In
        </Button>
        <Button className="w-[70%]">Create</Button>
      </CardFooter>
    </Card>
  );
}
