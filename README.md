# Fullstack
An example of packaging a Typescript frontend and Go backend together into a single binary.

In making this repo, I was specifically reproducing a minimalistic template of a pattern I saw in the Grafana repository. I thought it would be interesting to maintain a template of what I like to use.

Over time, this repo might become populated with some bloat for those that are minimalist. You may be interested in reviewing the commit history from before the creation of this README.md. I've attempted to take care in annotating commits with what happened in the repo to reach what I felt was a minimal template to start projects from.

## Examples Implemented
- Auth/Users

## Project Structure
This project's structure is dead fucking simple. All directories attempt to be WYSIWYG. If it is labeled "backend", that's probably the backend. There are some concepts or ideas you just have to be aware of in this project's ecosystem.

The easiest is to build the whole project and run it in a container.
```shell
make compose-up # Control+C to stop
make compose-down # Purges containers and data.
```

The prior method doesn't rely on any cached modules and is slow for iterating.

If you would like to iterate fast, open 3 terminal tabs and a code editor fixed on the root directory of this project.

First terminal, start the data services for development.
```shell
make services-up # Control+C to stop
make services-down # Purges containers and data.

# These are services you might want to access during development.
docker exec -it fullstack-yugabyte-1 ysqlsh -h yugabyte
```

Second terminal, start the backend. You will CTRL+C this and restart it for backend changes. 
```shell
make run-backend
```

Third terminal, start the frontend. You will navigate this in your browser. Vite server is configured to proxy `/api` to the backend you have in the last terminal.
```shell
make run-frontend
```

At this point if you're working on a component or part of the frontend, that will auto refresh in the browser. You still have to partly touch the terminal for the backend changes to reflect, but this is guaranteed to be quick to iterate on.


### Frontend
A lot of magic is happening in this directory because my plan is to offload all frontend work to GitHub Copilot. I don't want to be a frontend dev, but I've found that with the right tech stack available, GitHub Copilot is kind of a champ.

This repo is shipping:
- Tailwindcss
- Vite+React
- `frontend/frontend.go` is the laziest integration I could imagine with the `frontend/src/App.tsx` router.

I don't like the pathing for this, but most of the frontend work should probably happen in `frontend/src/components`.

Read up about [React](https://react.dev/).

The frontend must be built before running the go project!

```shell
npm run build
```

### Backend
Just use the utils to serve the responses. Offload persistence to alternate tooling. Most of the backend code should feel like workflows.

### Models
All the database operations are abstracted by GORM, because it enables some velocity in the development process. Yugabyte is my target database of choice, so the example model is based on some assumptions I have about how Yugabyte works.

Yugabyte is a distributed and replicated Postgres-compatible database. This is sorta the holy grail for SQL databases in the cloud. I don't know why there isn't more hype for the product. They maintain a somewhat decent helm chart. I like that in a potential vendor.

Due to the nature of being distributed, I suspect incremental IDs might be challenging at scale. Most problems are fine with a somewhat random distribution. I'm assuming primary keys as UUIDs in Yugabyte should be dispersed at scale.

A challenge with GORM is that you have to do some extra work to support UUIDs. This has been patterned out. Just follow the pattern and justify auto-increment operations.

If optimizations are necessary, a really good candidate for upgrades would be `github.com/jackc/pgx/v4/pgxpool`.

## Service Choices
These are some service choices I've been tinkering with. They provide different utility in making a distributed service.

### Metadata Service
Just use etcd clusters. I wrote a guide for deploying them as StatefulSets [here](https://etcd.io/docs/v3.5/op-guide/kubernetes/). I'm waiting to see what happens with this [operator here](https://github.com/etcd-io/etcd-operator).

There is kind of a lot you can do with these clusters. I think ETCD is a good candidate for maintaining consistency on user records in a distributed system. It could also be used like a filer.

Knowing about the reconciler pattern is useful for working with ETCD. The reconciler pattern is a way of ensuring that the state of a system is consistent with the desired state. This is useful for maintaining consistency in a distributed system.

### File Storage
Most social apps can rely on any S3-compatible object storage service that provides a CDN. Digital Ocean and Linode have some predictable pricing models, and the risk is pretty low given the following parameters:

- A file is being stored.
- Generate a random hex value (it should look like a sha512sum).
- Store the file at `/media/{random_sha512sum-like-value}.{original_ext}`.
- The file must be searchable somehow in `models/`.

This is fine for non-sensitive information like a social media site. 

For secure document exchange, access control might need to be manageable. This might be doable with CephFS's S3 implementation, or by just using a bigger cloud provider.

Regarding clients, [minio-go](https://github.com/minio/minio-go) is a decent pick for S3-compatible object storage.

### Alternative Databases
Dgraph looks interesting for a GraphQL database, but I'm questioning its consistency guarantees. They provide a decent helm chart for testing.

## Why GitHub Copilot is a Champ
I am not a frontend developer. However, Copilot/GPT-4o is proving to be a champ in agent mode with the tech stack chosen. I feel pretty comfortable letting the AI fuck up the frontend as long as I get to iteratively tell it "move pretty pixel left". So far it kinda can do that.

What this means though is great care must be made to keep things safe in the backend. 

- Anything generated by a user must be sanitized before a write. The requirements will vary depending on input restrictions. (Think along the lines: "Can this result in XSS?")
- It is easier to not hold sensitive information than it is to maintain an email in the database. Good apps should be making data useless in the event of a data leak.
- Establishing and documenting repeatable patterns will be helpful for reducing risk.
- Some of the risk can be mitigated by clever service choices. I want to stick to things I can operate in a Kubernetes cluster.

# Why AGPLv3
This is just the license I'm choosing for my projects. You replace that with whatever your project needs. 

# Contributing...
Don't. Just fork and own your own destiny.