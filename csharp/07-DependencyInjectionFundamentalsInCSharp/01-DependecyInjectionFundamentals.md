# Depenedency Injection Fundamentals in C# - DI vs IoC vs DIP

Starting example (Bad):

```csharp
var stars = new GithubService().GetStart("wiki");

// the higher level component
class GithubService
{
    public int GetStars(string repoName)
    {
        // here we rely on a smaller component (not good)
        return GithubClient().GetRepo(repoName).Stars;
    }
}

// the lower level component
class GithubClient
{
    public (string repoName, int Stars) GetRepo(string repoName)
    {
        return (repoName, repoName.Length);
    }
}
```

## Dependency Inversion Principle - DIP

Higher level components should not depend on lower level components.
Instead they should rely on abstractions.

Inversion of dependency example (Good):

```csharp
var stars = new GithubService().GetStars("wiki");

// the higher level component
class GithubService
{
    private IGithubClient _githubClient; // instanciating is still missing (see IoC)

    public int GetStars(string repoName)
    {
        // now we "rely" on an interface, instead of a class (good)
        return _githubClient.GetRepo(repoName).Stars;
    }
}

// the lower level component
class GithubClient : IGithubClient
{
    public (string repoName, int Stars) GetRepo(string repoName)
    {
        return (repoName, repoName.Length);
    }
}

internal interface IGithubClient
{
    (string repoName, int Stars) GetRepo(string repoName)
}
```

## Inversion of Control - IoC

Inversion of Control is a design principle in which you give the responsibility of the system to some container or a framework.

Three types of IoC:

1. Constructor injection

```csharp
//                  GithubClient has to be passed
var stars = new GithubService(<!!!>).GetStars("wiki");

//                  constructor inejciton
class GithubService(IGithubClient githubClient)
{
    private IGithubClient _githubClient = githubClient; 

    public int GetStars(string repoName)
    {
        return _githubClient.GetRepo(repoName).Stars;
    }
}
...
```

2. Setter injection

```csharp
var stars = new GithubService().GetStars("wiki");

class GithubService
{
    private IGithubClient _githubClient; 

    // setter injection
    public void SetGithubClient(IGithubClient githubClient)
    {
        _githubClient = githubClient; 
    }

    public int GetStars(string repoName)
    {
        return _githubClient.GetRepo(repoName).Stars;
    }
}
...
```

3. Interface injection

```csharp
var stars = new GithubService().GetStars("wiki");

//                    interface injection
class GithubService : IGithubClientSetter
{
    private IGithubClient _githubClient; 

    public void SetGithubClient(IGithubClient githubClient)
    {
        _githubClient = githubClient; 
    }

    public int GetStars(string repoName)
    {
        return _githubClient.GetRepo(repoName).Stars;
    }
}

internale interface IGithubClientSetter
{
    void SetGithubClient(IGithubClient githubClient)
}
...
```
