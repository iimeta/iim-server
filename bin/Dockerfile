FROM loads/alpine:3.8

###############################################################################
#                                 INSTALLATION                                #
###############################################################################

ENV WORKDIR        /app
ADD bin/iim-server $WORKDIR/bin/iim-server
ADD log            $WORKDIR/log
ADD manifest       $WORKDIR/manifest
RUN chmod +x       $WORKDIR/bin/iim-server

###############################################################################
#                                    START                                    #
###############################################################################

WORKDIR $WORKDIR
CMD ./bin/iim-server
